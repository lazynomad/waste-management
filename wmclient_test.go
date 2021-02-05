package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type mockHTTPClient struct {
}
func (client *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	testStr := `{
		"statusCode":200, 
		"data": {
			"access_token":"test.token.sign", 
			"id":"1a2b3c4d5e6f"
			}}`
	resp := &http.Response{
		StatusCode: 200,
		Body: ioutil.NopCloser(bytes.NewBufferString(testStr)),
	}

	return resp, nil
}

func TestGetAuthToken(t *testing.T) {
	testClient := getTestWmclient()
	token, _ := testClient.GetAuthToken()
	if (token != "test.token.sign") {
		t.Errorf("Wrong token" + token)
	}
}

func TestJsonMarshall(t *testing.T) {
	type authresp struct {
		StatusCode int `json:"statusCode"`
		Data struct {
			AccessToken string `json:"access_token"`
			ID string `json:"id"`
		} `json:"data"`
	}

	testStr := `{
		"statusCode":200, 
		"data": {
			"access_token":"test.token.sign", 
			"id":"1a2b3c4d5e6f"
			}}`
	body := ioutil.NopCloser(bytes.NewBufferString(testStr))
	resbody, _ := ioutil.ReadAll(body)
	//respbody := []byte(testStr)
	var p authresp
	err := json.Unmarshal(resbody, &p)
	if (err != nil) {
		panic(err.Error())
	}

	fmt.Println(string(p.Data.AccessToken))
}

func getTestWmclient() *WMClient {
	conf := Config {
		BaseURL: "https://test.url",
	}
	conf.Auth.User = "user"
	conf.Auth.Pass = "pass"
	conf.APIKeys.Auth = "1234567890"
	conf.APIKeys.Account = "1234567890"
	conf.APIKeys.Service = "1234567890"

	HTTPClient := new(mockHTTPClient)
	restClient := NewRestClient(HTTPClient)

	return NewWmClient(conf, *restClient)
}