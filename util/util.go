package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func Get2Response(url string, resp interface{}) error {
	result, err := http.DefaultClient.Get(url)
	if err != nil {
		return err
	}
	response, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return err
	}
	return nil
}

func Post2Response(url string, body io.Reader, resp interface{}) error {
	result, err := http.DefaultClient.Post(url, "application/json", body)
	if err != nil {
		return err
	}
	response, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return err
	}
	return nil
}
