package apps

import (
	"culfi/auth"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const apiURL string = "https://api.system.aws-usw02-pr.ice.predix.io"

// Apps contains App
type Apps struct {
	Apps []App `json:"resources"`
}

// App details container
type App struct {
	Name     string
	Memory   float64
	GUID     string
	Instaces float64
	State    string
}

// GetApps fetches all the apps
func GetApps() (Apps, error) {

	var apps Apps
	authToken, err := auth.ReadAuth()
	if err != nil {
		return apps, err
	}

	client := http.Client{}
	req, _ := http.NewRequest("GET", apiURL+"/v2/apps", nil)
	req.Header.Set("Authorization", authToken)
	res, _ := client.Do(req)
	respBody, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(respBody, &apps)
	if err != nil {
		return apps, err
	}

	return apps, nil
}

// UnmarshalJSON overrides the default Unmarshalling
func (app *App) UnmarshalJSON(b []byte) error {
	var f interface{}
	json.Unmarshal(b, &f)

	m := f.(map[string]interface{})

	entity := m["entity"]
	e := entity.(map[string]interface{})

	app.Name = e["name"].(string)
	app.Memory = e["memory"].(float64)
	app.Instaces = e["instances"].(float64)
	app.State = e["state"].(string)

	metadata := m["metadata"]
	md := metadata.(map[string]interface{})
	app.GUID = md["guid"].(string)

	return nil
}
