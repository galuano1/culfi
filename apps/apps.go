package apps

import (
	"culfi/auth"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const apiURL string = "https://api.system.aws-usw02-pr.ice.predix.io"

type apps struct {
	Apps []App `json:"apps"`
}

// App details container
type App struct {
	Name     string   `json:"name"`
	Memory   float64  `json:"memory"`
	GUID     string   `json:"guid"`
	Instaces float64  `json:"instances"`
	State    string   `json:"state"`
	URL      []string `json:"urls"`
}

// GetApps fetches all the apps
func GetApps(GUID string) ([]App, error) {

	var apps apps
	authToken, err := auth.ReadAuth()
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	req, _ := http.NewRequest("GET", apiURL+"/v2/spaces/"+GUID+"/summary", nil)
	req.Header.Set("Authorization", authToken)
	res, _ := client.Do(req)
	respBody, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(respBody, &apps)
	if err != nil {
		return nil, err
	}

	return apps.Apps, nil
}
