package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

//GET  http get method
func GET(baseURL string, p Parameters) ([]byte, error) {
	URL, _ := EncodeURL(baseURL, p)
	resp, _ := http.Get(URL)
	respBytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return respBytes, err
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
