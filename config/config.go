/*
Deals with loading and saving credentials and tokens to the config file.
*/
package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/shibukawa/configdir"
	keyring "github.com/zalando/go-keyring"

	"github.com/fstanis/go-hive/hive"
)

const (
	appName        = "hivecli"
	keyringService = appName

	defaultConfigFileName = "hivecli.conf"
)

var (
	ErrNoUsername = errors.New("no username stored in config")
	ErrNoURL      = errors.New("no URL stored in config")

	defaultConfigFilePath string
)

func init() {
	dirs := configdir.New("", appName).QueryFolders(configdir.Global)
	if len(dirs) > 0 {
		dirs[0].MkdirAll()
		defaultConfigFilePath = filepath.Join(dirs[0].Path, defaultConfigFileName)
	}
}

// Config represents the config file. This object will be serialized to the
// config file.
type Config struct {
	Username    string
	LoginURL    string
	EndpointURL string
	Token       string
}

// FromFile loads the config from the file with the given name.
func FromFile(filename string) (*Config, error) {
	if filename == "" {
		filename = defaultConfigFilePath
	}

	var conf Config

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return &conf, err
	}

	if err := json.Unmarshal(data, &conf); err != nil {
		return &conf, err
	}

	return &conf, nil
}

// ToFile saves the given config to the file with the given name.
func ToFile(conf *Config, filename string) error {
	if filename == "" {
		filename = defaultConfigFilePath
	}

	data, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0600)
}

// LoadCredentils constructs the hive.Credentials object using the data from the
// given Config combined with the password stored in the (platform dependent)
// keyring.
func (c *Config) LoadCredentials() (*hive.Credentials, error) {
	if c.Username == "" {
		return nil, ErrNoUsername
	}
	if c.LoginURL == "" {
		return nil, ErrNoURL
	}

	pass, err := keyring.Get(keyringService, c.Username)
	if err != nil {
		return nil, err
	}

	return &hive.Credentials{
		Username: c.Username,
		Password: pass,
		URL:      c.LoginURL,
	}, nil
}

// SaveCredentials saves the data from the given hive.Credentials object to the
// specified config while also storing the password in the (platform dependent)
// keyring.
func (c *Config) SaveCredentials(creds *hive.Credentials) error {
	c.Username = creds.Username
	c.LoginURL = creds.URL
	return keyring.Set(keyringService, creds.Username, creds.Password)
}
