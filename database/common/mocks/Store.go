// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"
	params "garm/params"

	mock "github.com/stretchr/testify/mock"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// AddInstanceEvent provides a mock function with given fields: ctx, instanceID, event, statusMessage
func (_m *Store) AddInstanceEvent(ctx context.Context, instanceID string, event params.EventType, statusMessage string) error {
	ret := _m.Called(ctx, instanceID, event, statusMessage)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, params.EventType, string) error); ok {
		r0 = rf(ctx, instanceID, event, statusMessage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ControllerInfo provides a mock function with given fields:
func (_m *Store) ControllerInfo() (params.ControllerInfo, error) {
	ret := _m.Called()

	var r0 params.ControllerInfo
	if rf, ok := ret.Get(0).(func() params.ControllerInfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(params.ControllerInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEnterprise provides a mock function with given fields: ctx, name, credentialsName, webhookSecret
func (_m *Store) CreateEnterprise(ctx context.Context, name string, credentialsName string, webhookSecret string) (params.Enterprise, error) {
	ret := _m.Called(ctx, name, credentialsName, webhookSecret)

	var r0 params.Enterprise
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) params.Enterprise); ok {
		r0 = rf(ctx, name, credentialsName, webhookSecret)
	} else {
		r0 = ret.Get(0).(params.Enterprise)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, name, credentialsName, webhookSecret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEnterprisePool provides a mock function with given fields: ctx, enterpriseID, param
func (_m *Store) CreateEnterprisePool(ctx context.Context, enterpriseID string, param params.CreatePoolParams) (params.Pool, error) {
	ret := _m.Called(ctx, enterpriseID, param)

	var r0 params.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string, params.CreatePoolParams) params.Pool); ok {
		r0 = rf(ctx, enterpriseID, param)
	} else {
		r0 = ret.Get(0).(params.Pool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, params.CreatePoolParams) error); ok {
		r1 = rf(ctx, enterpriseID, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateInstance provides a mock function with given fields: ctx, poolID, param
func (_m *Store) CreateInstance(ctx context.Context, poolID string, param params.CreateInstanceParams) (params.Instance, error) {
	ret := _m.Called(ctx, poolID, param)

	var r0 params.Instance
	if rf, ok := ret.Get(0).(func(context.Context, string, params.CreateInstanceParams) params.Instance); ok {
		r0 = rf(ctx, poolID, param)
	} else {
		r0 = ret.Get(0).(params.Instance)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, params.CreateInstanceParams) error); ok {
		r1 = rf(ctx, poolID, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrganization provides a mock function with given fields: ctx, name, credentialsName, webhookSecret
func (_m *Store) CreateOrganization(ctx context.Context, name string, credentialsName string, webhookSecret string) (params.Organization, error) {
	ret := _m.Called(ctx, name, credentialsName, webhookSecret)

	var r0 params.Organization
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) params.Organization); ok {
		r0 = rf(ctx, name, credentialsName, webhookSecret)
	} else {
		r0 = ret.Get(0).(params.Organization)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, name, credentialsName, webhookSecret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrganizationPool provides a mock function with given fields: ctx, orgId, param
func (_m *Store) CreateOrganizationPool(ctx context.Context, orgId string, param params.CreatePoolParams) (params.Pool, error) {
	ret := _m.Called(ctx, orgId, param)

	var r0 params.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string, params.CreatePoolParams) params.Pool); ok {
		r0 = rf(ctx, orgId, param)
	} else {
		r0 = ret.Get(0).(params.Pool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, params.CreatePoolParams) error); ok {
		r1 = rf(ctx, orgId, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRepository provides a mock function with given fields: ctx, owner, name, credentialsName, webhookSecret
func (_m *Store) CreateRepository(ctx context.Context, owner string, name string, credentialsName string, webhookSecret string) (params.Repository, error) {
	ret := _m.Called(ctx, owner, name, credentialsName, webhookSecret)

	var r0 params.Repository
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) params.Repository); ok {
		r0 = rf(ctx, owner, name, credentialsName, webhookSecret)
	} else {
		r0 = ret.Get(0).(params.Repository)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, owner, name, credentialsName, webhookSecret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRepositoryPool provides a mock function with given fields: ctx, repoId, param
func (_m *Store) CreateRepositoryPool(ctx context.Context, repoId string, param params.CreatePoolParams) (params.Pool, error) {
	ret := _m.Called(ctx, repoId, param)

	var r0 params.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string, params.CreatePoolParams) params.Pool); ok {
		r0 = rf(ctx, repoId, param)
	} else {
		r0 = ret.Get(0).(params.Pool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, params.CreatePoolParams) error); ok {
		r1 = rf(ctx, repoId, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *Store) CreateUser(ctx context.Context, user params.NewUserParams) (params.User, error) {
	ret := _m.Called(ctx, user)

	var r0 params.User
	if rf, ok := ret.Get(0).(func(context.Context, params.NewUserParams) params.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(params.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, params.NewUserParams) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteEnterprise provides a mock function with given fields: ctx, enterpriseID
func (_m *Store) DeleteEnterprise(ctx context.Context, enterpriseID string) error {
	ret := _m.Called(ctx, enterpriseID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, enterpriseID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteEnterprisePool provides a mock function with given fields: ctx, enterpriseID, poolID
func (_m *Store) DeleteEnterprisePool(ctx context.Context, enterpriseID string, poolID string) error {
	ret := _m.Called(ctx, enterpriseID, poolID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, enterpriseID, poolID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteInstance provides a mock function with given fields: ctx, poolID, instanceName
func (_m *Store) DeleteInstance(ctx context.Context, poolID string, instanceName string) error {
	ret := _m.Called(ctx, poolID, instanceName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, poolID, instanceName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOrganization provides a mock function with given fields: ctx, orgID
func (_m *Store) DeleteOrganization(ctx context.Context, orgID string) error {
	ret := _m.Called(ctx, orgID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, orgID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOrganizationPool provides a mock function with given fields: ctx, orgID, poolID
func (_m *Store) DeleteOrganizationPool(ctx context.Context, orgID string, poolID string) error {
	ret := _m.Called(ctx, orgID, poolID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, orgID, poolID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePoolByID provides a mock function with given fields: ctx, poolID
func (_m *Store) DeletePoolByID(ctx context.Context, poolID string) error {
	ret := _m.Called(ctx, poolID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, poolID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRepository provides a mock function with given fields: ctx, repoID
func (_m *Store) DeleteRepository(ctx context.Context, repoID string) error {
	ret := _m.Called(ctx, repoID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, repoID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRepositoryPool provides a mock function with given fields: ctx, repoID, poolID
func (_m *Store) DeleteRepositoryPool(ctx context.Context, repoID string, poolID string) error {
	ret := _m.Called(ctx, repoID, poolID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, repoID, poolID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindEnterprisePoolByTags provides a mock function with given fields: ctx, enterpriseID, tags
func (_m *Store) FindEnterprisePoolByTags(ctx context.Context, enterpriseID string, tags []string) (params.Pool, error) {
	ret := _m.Called(ctx, enterpriseID, tags)

	var r0 params.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) params.Pool); ok {
		r0 = rf(ctx, enterpriseID, tags)
	} else {
		r0 = ret.Get(0).(params.Pool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, enterpriseID, tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOrganizationPoolByTags provides a mock function with given fields: ctx, orgID, tags
func (_m *Store) FindOrganizationPoolByTags(ctx context.Context, orgID string, tags []string) (params.Pool, error) {
	ret := _m.Called(ctx, orgID, tags)

	var r0 params.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) params.Pool); ok {
		r0 = rf(ctx, orgID, tags)
	} else {
		r0 = ret.Get(0).(params.Pool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, orgID, tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindRepositoryPoolByTags provides a mock function with given fields: ctx, repoID, tags
func (_m *Store) FindRepositoryPoolByTags(ctx context.Context, repoID string, tags []string) (params.Pool, error) {
	ret := _m.Called(ctx, repoID, tags)

	var r0 params.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) params.Pool); ok {
		r0 = rf(ctx, repoID, tags)
	} else {
		r0 = ret.Get(0).(params.Pool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, repoID, tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEnterprise provides a mock function with given fields: ctx, name
func (_m *Store) GetEnterprise(ctx context.Context, name string) (params.Enterprise, error) {
	ret := _m.Called(ctx, name)

	var r0 params.Enterprise
	if rf, ok := ret.Get(0).(func(context.Context, string) params.Enterprise); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(params.Enterprise)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEnterpriseByID provides a mock function with given fields: ctx, enterpriseID
func (_m *Store) GetEnterpriseByID(ctx context.Context, enterpriseID string) (params.Enterprise, error) {
	ret := _m.Called(ctx, enterpriseID)

	var r0 params.Enterprise
	if rf, ok := ret.Get(0).(func(context.Context, string) params.Enterprise); ok {
		r0 = rf(ctx, enterpriseID)
	} else {
		r0 = ret.Get(0).(params.Enterprise)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, enterpriseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEnterprisePool provides a mock function with given fields: ctx, enterpriseID, poolID
func (_m *Store) GetEnterprisePool(ctx context.Context, enterpriseID string, poolID string) (params.Pool, error) {
	ret := _m.Called(ctx, enterpriseID, poolID)

	var r0 params.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) params.Pool); ok {
		r0 = rf(ctx, enterpriseID, poolID)
	} else {
		r0 = ret.Get(0).(params.Pool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, enterpriseID, poolID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInstanceByName provides a mock function with given fields: ctx, instanceName
func (_m *Store) GetInstanceByName(ctx context.Context, instanceName string) (params.Instance, error) {
	ret := _m.Called(ctx, instanceName)

	var r0 params.Instance
	if rf, ok := ret.Get(0).(func(context.Context, string) params.Instance); ok {
		r0 = rf(ctx, instanceName)
	} else {
		r0 = ret.Get(0).(params.Instance)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, instanceName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrganization provides a mock function with given fields: ctx, name
func (_m *Store) GetOrganization(ctx context.Context, name string) (params.Organization, error) {
	ret := _m.Called(ctx, name)

	var r0 params.Organization
	if rf, ok := ret.Get(0).(func(context.Context, string) params.Organization); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(params.Organization)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrganizationByID provides a mock function with given fields: ctx, orgID
func (_m *Store) GetOrganizationByID(ctx context.Context, orgID string) (params.Organization, error) {
	ret := _m.Called(ctx, orgID)

	var r0 params.Organization
	if rf, ok := ret.Get(0).(func(context.Context, string) params.Organization); ok {
		r0 = rf(ctx, orgID)
	} else {
		r0 = ret.Get(0).(params.Organization)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orgID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrganizationPool provides a mock function with given fields: ctx, orgID, poolID
func (_m *Store) GetOrganizationPool(ctx context.Context, orgID string, poolID string) (params.Pool, error) {
	ret := _m.Called(ctx, orgID, poolID)

	var r0 params.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) params.Pool); ok {
		r0 = rf(ctx, orgID, poolID)
	} else {
		r0 = ret.Get(0).(params.Pool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, orgID, poolID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPoolByID provides a mock function with given fields: ctx, poolID
func (_m *Store) GetPoolByID(ctx context.Context, poolID string) (params.Pool, error) {
	ret := _m.Called(ctx, poolID)

	var r0 params.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string) params.Pool); ok {
		r0 = rf(ctx, poolID)
	} else {
		r0 = ret.Get(0).(params.Pool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, poolID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPoolInstanceByName provides a mock function with given fields: ctx, poolID, instanceName
func (_m *Store) GetPoolInstanceByName(ctx context.Context, poolID string, instanceName string) (params.Instance, error) {
	ret := _m.Called(ctx, poolID, instanceName)

	var r0 params.Instance
	if rf, ok := ret.Get(0).(func(context.Context, string, string) params.Instance); ok {
		r0 = rf(ctx, poolID, instanceName)
	} else {
		r0 = ret.Get(0).(params.Instance)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, poolID, instanceName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRepository provides a mock function with given fields: ctx, owner, name
func (_m *Store) GetRepository(ctx context.Context, owner string, name string) (params.Repository, error) {
	ret := _m.Called(ctx, owner, name)

	var r0 params.Repository
	if rf, ok := ret.Get(0).(func(context.Context, string, string) params.Repository); ok {
		r0 = rf(ctx, owner, name)
	} else {
		r0 = ret.Get(0).(params.Repository)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, owner, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRepositoryByID provides a mock function with given fields: ctx, repoID
func (_m *Store) GetRepositoryByID(ctx context.Context, repoID string) (params.Repository, error) {
	ret := _m.Called(ctx, repoID)

	var r0 params.Repository
	if rf, ok := ret.Get(0).(func(context.Context, string) params.Repository); ok {
		r0 = rf(ctx, repoID)
	} else {
		r0 = ret.Get(0).(params.Repository)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, repoID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRepositoryPool provides a mock function with given fields: ctx, repoID, poolID
func (_m *Store) GetRepositoryPool(ctx context.Context, repoID string, poolID string) (params.Pool, error) {
	ret := _m.Called(ctx, repoID, poolID)

	var r0 params.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) params.Pool); ok {
		r0 = rf(ctx, repoID, poolID)
	} else {
		r0 = ret.Get(0).(params.Pool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, repoID, poolID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, user
func (_m *Store) GetUser(ctx context.Context, user string) (params.User, error) {
	ret := _m.Called(ctx, user)

	var r0 params.User
	if rf, ok := ret.Get(0).(func(context.Context, string) params.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(params.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: ctx, userID
func (_m *Store) GetUserByID(ctx context.Context, userID string) (params.User, error) {
	ret := _m.Called(ctx, userID)

	var r0 params.User
	if rf, ok := ret.Get(0).(func(context.Context, string) params.User); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(params.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasAdminUser provides a mock function with given fields: ctx
func (_m *Store) HasAdminUser(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// InitController provides a mock function with given fields:
func (_m *Store) InitController() (params.ControllerInfo, error) {
	ret := _m.Called()

	var r0 params.ControllerInfo
	if rf, ok := ret.Get(0).(func() params.ControllerInfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(params.ControllerInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllInstances provides a mock function with given fields: ctx
func (_m *Store) ListAllInstances(ctx context.Context) ([]params.Instance, error) {
	ret := _m.Called(ctx)

	var r0 []params.Instance
	if rf, ok := ret.Get(0).(func(context.Context) []params.Instance); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]params.Instance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllPools provides a mock function with given fields: ctx
func (_m *Store) ListAllPools(ctx context.Context) ([]params.Pool, error) {
	ret := _m.Called(ctx)

	var r0 []params.Pool
	if rf, ok := ret.Get(0).(func(context.Context) []params.Pool); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]params.Pool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListEnterpriseInstances provides a mock function with given fields: ctx, enterpriseID
func (_m *Store) ListEnterpriseInstances(ctx context.Context, enterpriseID string) ([]params.Instance, error) {
	ret := _m.Called(ctx, enterpriseID)

	var r0 []params.Instance
	if rf, ok := ret.Get(0).(func(context.Context, string) []params.Instance); ok {
		r0 = rf(ctx, enterpriseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]params.Instance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, enterpriseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListEnterprisePools provides a mock function with given fields: ctx, enterpriseID
func (_m *Store) ListEnterprisePools(ctx context.Context, enterpriseID string) ([]params.Pool, error) {
	ret := _m.Called(ctx, enterpriseID)

	var r0 []params.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string) []params.Pool); ok {
		r0 = rf(ctx, enterpriseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]params.Pool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, enterpriseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListEnterprises provides a mock function with given fields: ctx
func (_m *Store) ListEnterprises(ctx context.Context) ([]params.Enterprise, error) {
	ret := _m.Called(ctx)

	var r0 []params.Enterprise
	if rf, ok := ret.Get(0).(func(context.Context) []params.Enterprise); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]params.Enterprise)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOrgInstances provides a mock function with given fields: ctx, orgID
func (_m *Store) ListOrgInstances(ctx context.Context, orgID string) ([]params.Instance, error) {
	ret := _m.Called(ctx, orgID)

	var r0 []params.Instance
	if rf, ok := ret.Get(0).(func(context.Context, string) []params.Instance); ok {
		r0 = rf(ctx, orgID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]params.Instance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orgID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOrgPools provides a mock function with given fields: ctx, orgID
func (_m *Store) ListOrgPools(ctx context.Context, orgID string) ([]params.Pool, error) {
	ret := _m.Called(ctx, orgID)

	var r0 []params.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string) []params.Pool); ok {
		r0 = rf(ctx, orgID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]params.Pool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orgID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOrganizations provides a mock function with given fields: ctx
func (_m *Store) ListOrganizations(ctx context.Context) ([]params.Organization, error) {
	ret := _m.Called(ctx)

	var r0 []params.Organization
	if rf, ok := ret.Get(0).(func(context.Context) []params.Organization); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]params.Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPoolInstances provides a mock function with given fields: ctx, poolID
func (_m *Store) ListPoolInstances(ctx context.Context, poolID string) ([]params.Instance, error) {
	ret := _m.Called(ctx, poolID)

	var r0 []params.Instance
	if rf, ok := ret.Get(0).(func(context.Context, string) []params.Instance); ok {
		r0 = rf(ctx, poolID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]params.Instance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, poolID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRepoInstances provides a mock function with given fields: ctx, repoID
func (_m *Store) ListRepoInstances(ctx context.Context, repoID string) ([]params.Instance, error) {
	ret := _m.Called(ctx, repoID)

	var r0 []params.Instance
	if rf, ok := ret.Get(0).(func(context.Context, string) []params.Instance); ok {
		r0 = rf(ctx, repoID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]params.Instance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, repoID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRepoPools provides a mock function with given fields: ctx, repoID
func (_m *Store) ListRepoPools(ctx context.Context, repoID string) ([]params.Pool, error) {
	ret := _m.Called(ctx, repoID)

	var r0 []params.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string) []params.Pool); ok {
		r0 = rf(ctx, repoID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]params.Pool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, repoID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRepositories provides a mock function with given fields: ctx
func (_m *Store) ListRepositories(ctx context.Context) ([]params.Repository, error) {
	ret := _m.Called(ctx)

	var r0 []params.Repository
	if rf, ok := ret.Get(0).(func(context.Context) []params.Repository); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]params.Repository)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PoolInstanceCount provides a mock function with given fields: ctx, poolID
func (_m *Store) PoolInstanceCount(ctx context.Context, poolID string) (int64, error) {
	ret := _m.Called(ctx, poolID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, poolID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, poolID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEnterprise provides a mock function with given fields: ctx, enterpriseID, param
func (_m *Store) UpdateEnterprise(ctx context.Context, enterpriseID string, param params.UpdateRepositoryParams) (params.Enterprise, error) {
	ret := _m.Called(ctx, enterpriseID, param)

	var r0 params.Enterprise
	if rf, ok := ret.Get(0).(func(context.Context, string, params.UpdateRepositoryParams) params.Enterprise); ok {
		r0 = rf(ctx, enterpriseID, param)
	} else {
		r0 = ret.Get(0).(params.Enterprise)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, params.UpdateRepositoryParams) error); ok {
		r1 = rf(ctx, enterpriseID, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEnterprisePool provides a mock function with given fields: ctx, enterpriseID, poolID, param
func (_m *Store) UpdateEnterprisePool(ctx context.Context, enterpriseID string, poolID string, param params.UpdatePoolParams) (params.Pool, error) {
	ret := _m.Called(ctx, enterpriseID, poolID, param)

	var r0 params.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string, string, params.UpdatePoolParams) params.Pool); ok {
		r0 = rf(ctx, enterpriseID, poolID, param)
	} else {
		r0 = ret.Get(0).(params.Pool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, params.UpdatePoolParams) error); ok {
		r1 = rf(ctx, enterpriseID, poolID, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateInstance provides a mock function with given fields: ctx, instanceID, param
func (_m *Store) UpdateInstance(ctx context.Context, instanceID string, param params.UpdateInstanceParams) (params.Instance, error) {
	ret := _m.Called(ctx, instanceID, param)

	var r0 params.Instance
	if rf, ok := ret.Get(0).(func(context.Context, string, params.UpdateInstanceParams) params.Instance); ok {
		r0 = rf(ctx, instanceID, param)
	} else {
		r0 = ret.Get(0).(params.Instance)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, params.UpdateInstanceParams) error); ok {
		r1 = rf(ctx, instanceID, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOrganization provides a mock function with given fields: ctx, orgID, param
func (_m *Store) UpdateOrganization(ctx context.Context, orgID string, param params.UpdateRepositoryParams) (params.Organization, error) {
	ret := _m.Called(ctx, orgID, param)

	var r0 params.Organization
	if rf, ok := ret.Get(0).(func(context.Context, string, params.UpdateRepositoryParams) params.Organization); ok {
		r0 = rf(ctx, orgID, param)
	} else {
		r0 = ret.Get(0).(params.Organization)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, params.UpdateRepositoryParams) error); ok {
		r1 = rf(ctx, orgID, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOrganizationPool provides a mock function with given fields: ctx, orgID, poolID, param
func (_m *Store) UpdateOrganizationPool(ctx context.Context, orgID string, poolID string, param params.UpdatePoolParams) (params.Pool, error) {
	ret := _m.Called(ctx, orgID, poolID, param)

	var r0 params.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string, string, params.UpdatePoolParams) params.Pool); ok {
		r0 = rf(ctx, orgID, poolID, param)
	} else {
		r0 = ret.Get(0).(params.Pool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, params.UpdatePoolParams) error); ok {
		r1 = rf(ctx, orgID, poolID, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRepository provides a mock function with given fields: ctx, repoID, param
func (_m *Store) UpdateRepository(ctx context.Context, repoID string, param params.UpdateRepositoryParams) (params.Repository, error) {
	ret := _m.Called(ctx, repoID, param)

	var r0 params.Repository
	if rf, ok := ret.Get(0).(func(context.Context, string, params.UpdateRepositoryParams) params.Repository); ok {
		r0 = rf(ctx, repoID, param)
	} else {
		r0 = ret.Get(0).(params.Repository)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, params.UpdateRepositoryParams) error); ok {
		r1 = rf(ctx, repoID, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRepositoryPool provides a mock function with given fields: ctx, repoID, poolID, param
func (_m *Store) UpdateRepositoryPool(ctx context.Context, repoID string, poolID string, param params.UpdatePoolParams) (params.Pool, error) {
	ret := _m.Called(ctx, repoID, poolID, param)

	var r0 params.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string, string, params.UpdatePoolParams) params.Pool); ok {
		r0 = rf(ctx, repoID, poolID, param)
	} else {
		r0 = ret.Get(0).(params.Pool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, params.UpdatePoolParams) error); ok {
		r1 = rf(ctx, repoID, poolID, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, user, param
func (_m *Store) UpdateUser(ctx context.Context, user string, param params.UpdateUserParams) (params.User, error) {
	ret := _m.Called(ctx, user, param)

	var r0 params.User
	if rf, ok := ret.Get(0).(func(context.Context, string, params.UpdateUserParams) params.User); ok {
		r0 = rf(ctx, user, param)
	} else {
		r0 = ret.Get(0).(params.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, params.UpdateUserParams) error); ok {
		r1 = rf(ctx, user, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStore(t mockConstructorTestingTNewStore) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
