package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// GetAllPages will paginate through a page-operated URL until it encounters an empty string.
func GetAllPages(req *http.Request) (string, error) {
	var collatedPages []map[string]interface{}
	var returnBody []byte
	page := 1
	for string(returnBody) != "[]\n" { // while returned page is not empty
		q := req.URL.Query()
		q.Add("page", strconv.Itoa(page))
		req.URL.RawQuery = q.Encode()
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", err
		}
		err = failOnStatusCode(resp)
		if err != nil {
			return "", err
		}
		returnBody, err = getBody(resp)
		if err != nil {
			return "", err
		}
		collatedPages, err = combinePages(returnBody, collatedPages)
		if err != nil {
			return "", err
		}
		page++
	}
	collatedPagesJSON, err := json.Marshal(collatedPages)
	return string(collatedPagesJSON), err
}

func combinePages(addition []byte, toAddTo []map[string]interface{}) ([]map[string]interface{}, error) {
	var additionArray []map[string]interface{}
	err := json.Unmarshal([]byte(addition), &additionArray)
	if err != nil {
		return nil, err
	}
	fullArray := append(additionArray, toAddTo...)
	if err != nil {
		return nil, err
	}
	return fullArray, nil
}

func getBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
