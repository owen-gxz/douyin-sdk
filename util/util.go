package util

import (
	"encoding/json"
	"fmt"
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

func Post2Response2(url, token string, body io.Reader, resp interface{}) error {
	re, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return err
	}
	re.Header.Add("access-token", token)
	result, err := http.DefaultClient.Do(re)
	if err != nil {
		return err
	}
	response, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return err
	}
	return nil
}

func Get2Response2(url, token string, resp interface{}) error {
	re, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	re.Header.Add("access-token", token)
	result, err := http.DefaultClient.Do(re)
	if err != nil {
		return err
	}

	response, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return err
	}
	return nil
}
