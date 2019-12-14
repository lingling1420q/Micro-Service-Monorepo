package request

import (
	"bytes"
	"io/ioutil"
	"monaco/logger"
	"net/http"
)

//GET  http get method
func GET(baseURL string, p Parameters, header Parameters) ([]byte, error) {
	client := &http.Client{}
	URL, parameterErr := EncodeURL(baseURL, p)
	if parameterErr != nil {
		logger.Errorf("Parameter Error: %v", parameterErr)
		return []byte{}, parameterErr
	}

	req, _ := http.NewRequest("GET", URL, nil)
	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, mkClientErr := client.Do(req)
	if mkClientErr != nil {
		logger.Errorf("Make Http Client Error: %v", mkClientErr)
		return []byte{}, mkClientErr
	}
	defer resp.Body.Close()

	body, readBodyErr := ioutil.ReadAll(resp.Body)
	if readBodyErr != nil {
		logger.Errorf("Parse User Info Error: %v", readBodyErr)
		return []byte{}, readBodyErr
	}

	return body, nil
}

//POST  http post method
func POST(baseURL string, p Parameters, body []byte) ([]byte, error) {
	URL, err := EncodeURL(baseURL, p)
	resp, err := http.Post(URL, "application/json;charset=utf-8", bytes.NewBuffer(body))
	r, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return r, err
}

//DELETE  http delete method
func DELETE() {

}

//PUT  http put method
func PUT() {

}
