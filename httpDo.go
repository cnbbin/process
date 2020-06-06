package process

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func HttpJsonDo(url string, jsonMarshal []byte, httpMethod string) ([]byte, error) {
	var err error
	var response []byte
	client := &http.Client{}
	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(jsonMarshal))
	if err != nil {
		return response, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	return response, nil
}
