// Copyright 2024 Cloudbase Solutions SRL
//
//	Licensed under the Apache License, Version 2.0 (the "License"); you may
//	not use this file except in compliance with the License. You may obtain
//	a copy of the License at
//
//	     http://www.apache.org/licenses/LICENSE-2.0
//
//	Unless required by applicable law or agreed to in writing, software
//	distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//	WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//	License for the specific language governing permissions and limitations
//	under the License.

package sql

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	runnerErrors "github.com/cloudbase/garm-provider-common/errors"
	"github.com/cloudbase/garm-provider-common/util"
	"github.com/cloudbase/garm/params"
)

func (s *sqlDatabase) CreateEnterprise(ctx context.Context, name, credentialsName, webhookSecret string, poolBalancerType params.PoolBalancerType) (params.Enterprise, error) {
	if webhookSecret == "" {
		return params.Enterprise{}, errors.New("creating enterprise: missing secret")
	}
	secret, err := util.Seal([]byte(webhookSecret), []byte(s.cfg.Passphrase))
	if err != nil {
		return params.Enterprise{}, errors.Wrap(err, "encoding secret")
	}
	newEnterprise := Enterprise{
		Name:             name,
		WebhookSecret:    secret,
		CredentialsName:  credentialsName,
		PoolBalancerType: poolBalancerType,
	}
	err = s.conn.Transaction(func(tx *gorm.DB) error {
		creds, err := s.getGithubCredentialsByName(ctx, tx, credentialsName, false)
		if err != nil {
			return errors.Wrap(err, "creating enterprise")
		}
		newEnterprise.CredentialsID = &creds.ID

		q := tx.Create(&newEnterprise)
		if q.Error != nil {
			return errors.Wrap(q.Error, "creating enterprise")
		}

		newEnterprise.Credentials = creds

		return nil
	})
	if err != nil {
		return params.Enterprise{}, errors.Wrap(err, "creating enterprise")
	}

	param, err := s.sqlToCommonEnterprise(newEnterprise, true)
	if err != nil {
		return params.Enterprise{}, errors.Wrap(err, "creating enterprise")
	}

	return param, nil
}

func (s *sqlDatabase) GetEnterprise(ctx context.Context, name string) (params.Enterprise, error) {
	enterprise, err := s.getEnterprise(ctx, name)
	if err != nil {
		return params.Enterprise{}, errors.Wrap(err, "fetching enterprise")
	}

	param, err := s.sqlToCommonEnterprise(enterprise, true)
	if err != nil {
		return params.Enterprise{}, errors.Wrap(err, "fetching enterprise")
	}
	return param, nil
}

func (s *sqlDatabase) GetEnterpriseByID(ctx context.Context, enterpriseID string) (params.Enterprise, error) {
	enterprise, err := s.getEnterpriseByID(ctx, s.conn, enterpriseID, "Pools", "Credentials", "Endpoint")
	if err != nil {
		return params.Enterprise{}, errors.Wrap(err, "fetching enterprise")
	}

	param, err := s.sqlToCommonEnterprise(enterprise, true)
	if err != nil {
		return params.Enterprise{}, errors.Wrap(err, "fetching enterprise")
	}
	return param, nil
}

func (s *sqlDatabase) ListEnterprises(_ context.Context) ([]params.Enterprise, error) {
	var enterprises []Enterprise
	q := s.conn.Preload("Credentials").Find(&enterprises)
	if q.Error != nil {
		return []params.Enterprise{}, errors.Wrap(q.Error, "fetching enterprises")
	}

	ret := make([]params.Enterprise, len(enterprises))
	for idx, val := range enterprises {
		var err error
		ret[idx], err = s.sqlToCommonEnterprise(val, true)
		if err != nil {
			return nil, errors.Wrap(err, "fetching enterprises")
		}
	}

	return ret, nil
}

func (s *sqlDatabase) DeleteEnterprise(ctx context.Context, enterpriseID string) error {
	enterprise, err := s.getEnterpriseByID(ctx, s.conn, enterpriseID)
	if err != nil {
		return errors.Wrap(err, "fetching enterprise")
	}

	q := s.conn.Unscoped().Delete(&enterprise)
	if q.Error != nil && !errors.Is(q.Error, gorm.ErrRecordNotFound) {
		return errors.Wrap(q.Error, "deleting enterprise")
	}

	return nil
}

func (s *sqlDatabase) UpdateEnterprise(ctx context.Context, enterpriseID string, param params.UpdateEntityParams) (params.Enterprise, error) {
	var enterprise Enterprise
	var creds GithubCredentials
	err := s.conn.Transaction(func(tx *gorm.DB) error {
		var err error
		enterprise, err = s.getEnterpriseByID(ctx, tx, enterpriseID, "Credentials", "Endpoint")
		if err != nil {
			return errors.Wrap(err, "fetching enterprise")
		}

		if param.CredentialsName != "" {
			creds, err = s.getGithubCredentialsByName(ctx, tx, param.CredentialsName, false)
			if err != nil {
				return errors.Wrap(err, "fetching credentials")
			}
			enterprise.CredentialsID = &creds.ID
		}
		if param.WebhookSecret != "" {
			secret, err := util.Seal([]byte(param.WebhookSecret), []byte(s.cfg.Passphrase))
			if err != nil {
				return errors.Wrap(err, "encoding secret")
			}
			enterprise.WebhookSecret = secret
		}

		if param.PoolBalancerType != "" {
			enterprise.PoolBalancerType = param.PoolBalancerType
		}

		q := tx.Save(&enterprise)
		if q.Error != nil {
			return errors.Wrap(q.Error, "saving enterprise")
		}

		if creds.ID != 0 {
			enterprise.Credentials = creds
		}

		return nil
	})
	if err != nil {
		return params.Enterprise{}, errors.Wrap(err, "updating enterprise")
	}

	newParams, err := s.sqlToCommonEnterprise(enterprise, true)
	if err != nil {
		return params.Enterprise{}, errors.Wrap(err, "updating enterprise")
	}
	return newParams, nil
}

func (s *sqlDatabase) getEnterprise(_ context.Context, name string) (Enterprise, error) {
	var enterprise Enterprise

	q := s.conn.Where("name = ? COLLATE NOCASE", name).
		Preload("Credentials").
		Preload("Endpoint").
		First(&enterprise)
	if q.Error != nil {
		if errors.Is(q.Error, gorm.ErrRecordNotFound) {
			return Enterprise{}, runnerErrors.ErrNotFound
		}
		return Enterprise{}, errors.Wrap(q.Error, "fetching enterprise from database")
	}
	return enterprise, nil
}

func (s *sqlDatabase) getEnterpriseByID(_ context.Context, tx *gorm.DB, id string, preload ...string) (Enterprise, error) {
	u, err := uuid.Parse(id)
	if err != nil {
		return Enterprise{}, errors.Wrap(runnerErrors.ErrBadRequest, "parsing id")
	}
	var enterprise Enterprise

	q := tx
	if len(preload) > 0 {
		for _, field := range preload {
			q = q.Preload(field)
		}
	}
	q = q.Where("id = ?", u).First(&enterprise)

	if q.Error != nil {
		if errors.Is(q.Error, gorm.ErrRecordNotFound) {
			return Enterprise{}, runnerErrors.ErrNotFound
		}
		return Enterprise{}, errors.Wrap(q.Error, "fetching enterprise from database")
	}
	return enterprise, nil
}
