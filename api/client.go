package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

type buildkiteClient struct {
	apiToken     string
	organization string
}

func makeClient() *buildkiteClient {
	client := &buildkiteClient{}
	client.organization = viper.GetString("organization")
	client.apiToken = viper.GetString("token")
	return client
}

func buildBaseRequest(urlPath, method string) (*http.Request, error) {
	client := makeClient()
	request, err := http.NewRequest(method, fmt.Sprintf("https://api.buildkite.com/v2/organizations/%s/%s", client.organization, urlPath), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.apiToken))
	return request, nil
}

// GetAgentList will return JSON string of agent list
func GetAgentList() (string, error) {
	req, err := buildBaseRequest("agents", "GET")
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
