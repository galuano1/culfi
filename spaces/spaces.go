package spaces

import (
	"culfi/auth"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Space contains the cloud foundry space details
type Space struct {
	Name    string
	OrgName string
	GUID    string
}

type spaces struct {
	Spaces []Space `json:"resources"`
}

type org struct {
	Name string
}

const apiURL string = "https://api.system.aws-usw02-pr.ice.predix.io"

// GetSpaces gets the spaces available to the user
func GetSpaces() ([]Space, error) {

	authToken, err := auth.ReadAuth()
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	req, _ := http.NewRequest("GET", apiURL+"/v2/spaces", nil)
	req.Header.Set("Authorization", authToken)
	res, _ := client.Do(req)
	respBody, _ := ioutil.ReadAll(res.Body)

	var spaces spaces
	err = json.Unmarshal(respBody, &spaces)
	if err != nil {
		return nil, err
	}

	return spaces.Spaces, nil
}

// UnmarshalJSON overrides the default Unmarshalling
func (space *Space) UnmarshalJSON(b []byte) error {
	var f interface{}
	json.Unmarshal(b, &f)

	m := f.(map[string]interface{})

	entity := m["entity"]
	e := entity.(map[string]interface{})

	space.Name = e["name"].(string)
	orgURL := e["organization_url"].(string)
	space.OrgName, _ = getOrgName(orgURL)
	metadata := m["metadata"]
	md := metadata.(map[string]interface{})
	space.GUID = md["guid"].(string)

	return nil
}

func getOrgName(url string) (string, error) {
	orgURL := apiURL + url
	client := http.Client{}
	req, err := http.NewRequest("GET", orgURL, nil)
	if err != nil {
		return "", err
	}

	authToken, err := auth.ReadAuth()
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", authToken)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var f interface{}
	json.Unmarshal(respBody, &f)

	m := f.(map[string]interface{})
	entity := m["entity"]
	e := entity.(map[string]interface{})

	name := e["name"].(string)

	return name, nil
}
