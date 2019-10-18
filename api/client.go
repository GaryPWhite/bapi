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

func failOnStatusCode(res *http.Response) error {
	if res.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf("Failed on status code %d, Response:\n%s", res.StatusCode, b)
	}
	return nil
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
	agentsString, err := GetAllPages(req)
	if err != nil {
		return "", err
	}
	return agentsString, err
}

// GetBuildsList will return JSON string of all builds
func GetBuildsList() (string, error) {
	var err error
	var req *http.Request
	if viper.IsSet("pipeline") {
		req, err = buildBaseRequest(fmt.Sprintf("pipelines/%s/builds", viper.GetString("pipeline")), "GET")
	} else {
		req, err = buildBaseRequest("builds", "GET")
	}
	if err != nil {
		return "", err
	}
	agentsString, err := GetAllPages(req)
	if err != nil {
		return "", err
	}
	return agentsString, err
}

// StopAgent will stop specified agent
func StopAgent(agentID string, force bool) error {
	req, err := buildBaseRequest(fmt.Sprintf("/agents/%s/stop", agentID), "PUT")
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return failOnStatusCode(res)
}
