// Copyright 2022 Cloudbase Solutions SRL
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	EncryptionPassphrase     = "bocyasicgatEtenOubwonIbsudNutDom"
	WeakEncryptionPassphrase = "1234567890abcdefghijklmnopqrstuv"
)

func getDefaultSectionConfig(configDir string) Default {
	return Default{
		ConfigDir:   configDir,
		CallbackURL: "https://garm.example.com/",
		LogFile:     filepath.Join(configDir, "garm.log"),
	}
}

func getDefaultTLSConfig() TLSConfig {
	return TLSConfig{
		CRT:    "../testdata/certs/srv-pub.pem",
		Key:    "../testdata/certs/srv-key.pem",
		CACert: "../testdata/certs/ca-pub.pem",
	}
}

func getDefaultAPIServerConfig() APIServer {
	return APIServer{
		Bind:        "0.0.0.0",
		Port:        9998,
		UseTLS:      true,
		TLSConfig:   getDefaultTLSConfig(),
		CORSOrigins: []string{},
	}
}

func getMySQLDefaultConfig() MySQL {
	return MySQL{
		Username:     "test",
		Password:     "test",
		Hostname:     "127.0.0.1",
		DatabaseName: "garm",
	}
}

func getDefaultDatabaseConfig(dir string) Database {
	return Database{
		Debug:     false,
		DbBackend: SQLiteBackend,
		SQLite: SQLite{
			DBFile: filepath.Join(dir, "garm.db"),
		},
		Passphrase: EncryptionPassphrase,
	}
}

func getDefaultProvidersConfig() []Provider {
	lxdConfig := getDefaultLXDConfig()
	return []Provider{
		{
			Name:         "test_lxd",
			ProviderType: LXDProvider,
			Description:  "test LXD provider",
			LXD:          lxdConfig,
		},
	}
}

func getDefaultGithubConfig() []Github {
	return []Github{
		{
			Name:        "dummy_creds",
			Description: "dummy github credentials",
			OAuth2Token: "bogus",
		},
	}
}

func getDefaultJWTCofig() JWTAuth {
	return JWTAuth{
		Secret:     EncryptionPassphrase,
		TimeToLive: "48h",
	}
}

func getDefaultConfig(t *testing.T) Config {
	dir, err := ioutil.TempDir("", "garm-config-test")
	if err != nil {
		t.Fatalf("failed to create temporary directory: %s", err)
	}
	t.Cleanup(func() { os.RemoveAll(dir) })

	return Config{
		Default:   getDefaultSectionConfig(dir),
		APIServer: getDefaultAPIServerConfig(),
		Database:  getDefaultDatabaseConfig(dir),
		Providers: getDefaultProvidersConfig(),
		Github:    getDefaultGithubConfig(),
		JWTAuth:   getDefaultJWTCofig(),
	}
}

func TestConfig(t *testing.T) {
	cfg := getDefaultConfig(t)

	err := cfg.Validate()
	require.Nil(t, err)
}

func TestDefaultSectionConfig(t *testing.T) {
	dir, err := ioutil.TempDir("", "garm-config-test")
	if err != nil {
		t.Fatalf("failed to create temporary directory: %s", err)
	}
	t.Cleanup(func() { os.RemoveAll(dir) })
	cfg := getDefaultSectionConfig(dir)

	tests := []struct {
		name      string
		cfg       Default
		errString string
	}{
		{
			name:      "Config is valid",
			cfg:       cfg,
			errString: "",
		},
		{
			name: "CallbackURL cannot be empty",
			cfg: Default{
				CallbackURL: "",
				ConfigDir:   cfg.ConfigDir,
			},
			errString: "missing callback_url",
		},
		{
			name: "ConfigDir cannot be empty",
			cfg: Default{
				CallbackURL: cfg.CallbackURL,
				ConfigDir:   "",
			},
			errString: "config_dir cannot be empty",
		},
		{
			name: "config_dir must exist and be accessible",
			cfg: Default{
				CallbackURL: cfg.CallbackURL,
				ConfigDir:   "/i/do/not/exist",
			},
			errString: "accessing config dir: stat /i/do/not/exist:.*",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cfg.Validate()
			if tc.errString == "" {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
				require.Regexp(t, tc.errString, err.Error())
			}
		})
	}
}

func TestValidateAPIServerConfig(t *testing.T) {
	cfg := getDefaultAPIServerConfig()

	tests := []struct {
		name      string
		cfg       APIServer
		errString string
	}{
		{
			name:      "Config is valid",
			cfg:       cfg,
			errString: "",
		},
		{
			name: "Bind address is empty",
			cfg: APIServer{
				Bind: "",
				Port: 9998,
			},
			errString: "invalid IP address",
		},
		{
			name: "Bind address is invalid",
			cfg: APIServer{
				Bind: "not an IP",
				Port: 9998,
			},
			errString: "invalid IP address",
		},
		{
			name: "Bind address is valid IPv6",
			cfg: APIServer{
				Bind: "::",
				Port: 9998,
			},
			errString: "",
		},
		{
			name: "Port is not set",
			cfg: APIServer{
				Bind: cfg.Bind,
				Port: 0,
			},
			errString: "invalid port nr 0",
		},
		{
			name: "Port is not valid",
			cfg: APIServer{
				Bind: cfg.Bind,
				Port: 65536,
			},
			errString: "invalid port nr 65536",
		},
		{
			name: "Invalid TLS config",
			cfg: APIServer{
				Bind:      cfg.Bind,
				Port:      cfg.Port,
				TLSConfig: TLSConfig{},
				UseTLS:    true,
			},
			errString: "TLS validation failed:*",
		},
		{
			name: "Skip TLS config validation if UseTLS is false",
			cfg: APIServer{
				Bind:      cfg.Bind,
				Port:      cfg.Port,
				TLSConfig: TLSConfig{},
				UseTLS:    false,
			},
			errString: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cfg.Validate()
			if tc.errString == "" {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
				require.Regexp(t, tc.errString, err.Error())
			}
		})
	}
}

func TestAPIBindAddress(t *testing.T) {
	cfg := getDefaultAPIServerConfig()

	err := cfg.Validate()
	require.Nil(t, err)
	require.Equal(t, cfg.BindAddress(), "0.0.0.0:9998")
}

func TestAPITLSconfig(t *testing.T) {
	cfg := getDefaultAPIServerConfig()

	err := cfg.Validate()
	require.Nil(t, err)

	tlsCfg, err := cfg.APITLSConfig()
	require.Nil(t, err)
	require.NotNil(t, tlsCfg)

	// Any error in the TLSConfig should return an error here.
	cfg.TLSConfig = TLSConfig{}
	tlsCfg, err = cfg.APITLSConfig()
	require.NotNil(t, err)
	require.EqualError(t, err, "missing crt or key")

	// If TLS is disabled, don't validate TLSconfig.
	cfg.UseTLS = false
	tlsCfg, err = cfg.APITLSConfig()
	require.Nil(t, err)
	require.Nil(t, tlsCfg)
}

func TestTLSConfig(t *testing.T) {
	dir, err := ioutil.TempDir("", "garm-config-test")
	if err != nil {
		t.Fatalf("failed to create temporary directory: %s", err)
	}
	t.Cleanup(func() { os.RemoveAll(dir) })

	invalidCert := filepath.Join(dir, "invalid_cert.pem")
	err = ioutil.WriteFile(invalidCert, []byte("bogus content"), 0755)
	if err != nil {
		t.Fatalf("failed to write file: %s", err)
	}

	cfg := getDefaultTLSConfig()

	err = cfg.Validate()
	require.Nil(t, err)

	tests := []struct {
		name      string
		cfg       TLSConfig
		errString string
	}{
		{
			name:      "Config is valid",
			cfg:       cfg,
			errString: "",
		},
		{
			name: "missing crt",
			cfg: TLSConfig{
				CRT:    "",
				Key:    cfg.Key,
				CACert: cfg.CACert,
			},
			errString: "missing crt or key",
		},
		{
			name: "missing key",
			cfg: TLSConfig{
				CRT:    cfg.CRT,
				Key:    "",
				CACert: cfg.CACert,
			},
			errString: "missing crt or key",
		},
		{
			name: "invalid CA cert",
			cfg: TLSConfig{
				CRT:    cfg.CRT,
				Key:    cfg.Key,
				CACert: invalidCert,
			},
			errString: "failed to parse CA cert",
		},
		{
			name: "invalid cert",
			cfg: TLSConfig{
				CRT:    invalidCert,
				Key:    cfg.Key,
				CACert: cfg.CACert,
			},
			errString: "tls: failed to find any PEM data in certificate input",
		},
		{
			name: "invalid key",
			cfg: TLSConfig{
				CRT:    cfg.CRT,
				Key:    invalidCert,
				CACert: cfg.CACert,
			},
			errString: "tls: failed to find any PEM data in key input",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tlsCfg, err := tc.cfg.TLSConfig()
			if tc.errString == "" {
				require.Nil(t, err)
				require.NotNil(t, tlsCfg)
			} else {
				require.NotNil(t, err)
				require.Nil(t, tlsCfg)
				require.Regexp(t, tc.errString, err.Error())
			}
		})
	}
}

func TestDatabaseConfig(t *testing.T) {
	dir, err := ioutil.TempDir("", "garm-config-test")
	if err != nil {
		t.Fatalf("failed to create temporary directory: %s", err)
	}
	t.Cleanup(func() { os.RemoveAll(dir) })
	cfg := getDefaultDatabaseConfig(dir)

	tests := []struct {
		name      string
		cfg       Database
		errString string
	}{
		{
			name:      "Config is valid",
			cfg:       cfg,
			errString: "",
		},
		{
			name: "Missing backend",
			cfg: Database{
				DbBackend:  "",
				SQLite:     cfg.SQLite,
				Passphrase: cfg.Passphrase,
			},
			errString: "invalid databse configuration: backend is required",
		},
		{
			name: "Invalid backend type",
			cfg: Database{
				DbBackend:  DBBackendType("bogus"),
				SQLite:     cfg.SQLite,
				Passphrase: cfg.Passphrase,
			},
			errString: "invalid database backend: bogus",
		},
		{
			name: "Missing passphrase",
			cfg: Database{
				DbBackend:  cfg.DbBackend,
				SQLite:     cfg.SQLite,
				Passphrase: "",
			},
			errString: "passphrase must be set and it must be a string of 32 characters*",
		},
		{
			name: "passphrase has invalid length",
			cfg: Database{
				DbBackend:  cfg.DbBackend,
				SQLite:     cfg.SQLite,
				Passphrase: "testing",
			},
			errString: "passphrase must be set and it must be a string of 32 characters*",
		},
		{
			name: "passphrase is too weak",
			cfg: Database{
				DbBackend:  cfg.DbBackend,
				SQLite:     cfg.SQLite,
				Passphrase: WeakEncryptionPassphrase,
			},
			errString: "database passphrase is too weak",
		},
		{
			name: "sqlite3 backend is missconfigured",
			cfg: Database{
				DbBackend: cfg.DbBackend,
				SQLite: SQLite{
					DBFile: "",
				},
				Passphrase: cfg.Passphrase,
			},
			errString: "validating sqlite3 config: no valid db_file was specified",
		},
		{
			name: "mysql backend is missconfigured",
			cfg: Database{
				DbBackend:  MySQLBackend,
				MySQL:      MySQL{},
				Passphrase: cfg.Passphrase,
			},
			errString: "validating mysql config: database, username, password, hostname are mandatory parameters for the database section",
		},
		{
			name: "mysql backend is configured and valid",
			cfg: Database{
				DbBackend:  MySQLBackend,
				MySQL:      getMySQLDefaultConfig(),
				Passphrase: cfg.Passphrase,
			},
			errString: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cfg.Validate()
			if tc.errString == "" {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
				require.Regexp(t, tc.errString, err.Error())
			}
		})
	}
}
