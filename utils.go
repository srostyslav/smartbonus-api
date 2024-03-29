package smartbonus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var rootPath string                                           // route of smartbonus: ask smartbonus team for you
const testStoreId = "fa7cdd67-e973-4125-9990-9c50d61faa3a"    // test store id used for testing api
const testCustomerId = "0e803ff6-70c6-4c0e-93af-a7fd8b6352f5" // test store id used for testing api
const testUserId = "380662436090"                             // test user id used for testing api

type sbResponse struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}

// Helper function that decode smartbonus response to your object
func decodeJson(resp *http.Response, response []byte, obj interface{}) error {

	var sbResp *sbResponse
	if err := json.Unmarshal(response, &sbResp); err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		if message, err := json.Marshal(sbResp.Message); err != nil {
			return err
		} else {
			return json.Unmarshal(message, &obj)
		}
	} else {
		return fmt.Errorf("%v", sbResp.Message)
	}
}

// Helper function that send POST request encoding in json
func sendPostRequest(url string, body interface{}, obj interface{}) error {

	output, err := json.Marshal(&body)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "encoding/json", bytes.NewBuffer(output))
	if err != nil {
		return err
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return decodeJson(resp, response, obj)
}

// Helper function that send GET request encoding in json
func sendGetRequest(url string, params map[string]string, obj interface{}) error {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return decodeJson(resp, response, obj)
}
