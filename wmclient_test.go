package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/lazynomad/waste-management/restclient"
)

// Mock HTTP client that implements the interface defined in the restclient package.
type mockHTTPClient struct {
}

var (
	// Used to call mockHTTPClient Do function.
	// This function is used to define the test action when an HTTP call is executed.
	doFunc func(req *http.Request) (*http.Response, error)
)

// Implementation method of mockHTTPClient following the interface defined in restclient package.
func (client *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return doFunc(req)
}

// Tests the flow of actions to fetch the Auth token
func TestGetAuthToken(t *testing.T) {
	testClient := getTestWmclient()

	doFunc = func(*http.Request) (*http.Response, error) {
		testStr := `{
			"statusCode":200, 
			"data": {
				"access_token":"test.token.sign", 
				"id":"1a2b3c4d5e6f"
				}}`

		return &http.Response{
			StatusCode: 200,
			Body: ioutil.NopCloser(bytes.NewBufferString(testStr)),
		}, nil
	}

	token, err := testClient.GetAuthToken()

	if (err != nil) {
		t.Errorf("Failed with error" + err.Error())
	}

	if (token != "test.token.sign") {
		t.Errorf("Wrong token" + token)
	}
}

// Gets a wmclient stub with dummy configs over mock HTTP client
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
	restClient := restclient.NewRestClient(HTTPClient)

	return NewWmClient(conf, *restClient)
}