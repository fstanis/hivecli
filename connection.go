package main

import (
	"fmt"

	"github.com/fstanis/go-hive/hive"
	"github.com/fstanis/hivecli/config"
)

// Logs in using the specified `hive.Credentials` and, if successful, saves both
// the credentials and the received token and endpoint to the given config.
func login(conf *config.Config, creds *hive.Credentials) (*hive.Client, error) {
	client := hive.NewClient()
	if err := client.Login(creds); err != nil {
		return nil, err
	}
	conf.SaveCredentials(creds)
	conf.Token = client.Token
	conf.EndpointURL = client.EndpointURL
	config.ToFile(conf, *configPath)
	return client, nil
}

// Restores the token from the given config and attempts to use it.
func restoreToken(conf *config.Config) (*hive.Client, error) {
	client := hive.NewClient()
	client.Token = conf.Token
	client.EndpointURL = conf.EndpointURL
	err := client.RefreshDevices()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return client, nil
}

// Performs the following action:
// 1.  Tries to load the config file.
// 2.  If the config file contains no valid credentials, asks the user.
// 3.  If the config contains a token, attempts to use it.
// 4.  If no valid token can be used, logs in.
// 5.  Saves the credentials and token into the config.
func connect(conf *config.Config) (*hive.Client, error) {
	creds, err := conf.LoadCredentials()
	if err != nil {
		fmt.Println("No valid credentials found, please input them now.\n")
		creds = askCredentials()
	}

	if conf.Token != "" && conf.EndpointURL != "" {
		client, err := restoreToken(conf)
		if err == nil {
			return client, nil
		}
	}

	return login(conf, creds)
}
