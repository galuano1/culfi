package auth

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type config struct {
	AccessToken           string
	AuthorizationEndpoint string
}

// ReadAuth reads the authorization token from the cloud foundry config file
func ReadAuth() (string, error) {
	homeDir := os.Getenv("HOME")
	configFile := homeDir + "/.cf/config.json"

	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		return "", err
	}

	var myConfig config
	err = json.Unmarshal(content, &myConfig)
	if err != nil {
		return "", err
	}
	return myConfig.AccessToken, nil

}
