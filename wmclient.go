package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"path"
)

type authresp struct {
	StatusCode int `json:"statusCode"`
	Data struct {
		AccessToken string `json:"access_token"`
		ID string `json:"id"`
	} `json:"data"`
}

// WMClient to communicate with waste management endpoint
type WMClient struct {
	conf Config
	restClient RestClient
}

// GetAuthToken to get the JWT token from Waste-Management
func (wm *WMClient) GetAuthToken() (string, error) {
	fullURL := genURL(wm.conf.BaseURL, "user/authenticate")
	jsonVal := map[string]string{"username": wm.conf.Auth.User, "password": wm.conf.Auth.Pass, "locale": "en_US"}
	jsonM, _ := json.Marshal(jsonVal)
	headers := getDefaultHeaders()
	headers["apikey"] = wm.conf.APIKeys.Auth

	respCode, respBody, err := wm.restClient.Get(fullURL, jsonM, headers)
	if (err != nil) {
		panic(err.Error())
	}
	
	if (respCode != 200) {
		fmt.Printf("Failed. Code: %d. Body: %s", respCode, string(respBody))
		errMsg := fmt.Sprintf("Unable to fetch auth token. Response code: %d", respCode)
		return "", errors.New(errMsg)
	}

	parsed, err := parseAuthResp(respBody)
	if (err != nil) {
		panic(err.Error())
	}

	fmt.Printf("status: %d - id: %s - token:%s", parsed.StatusCode, parsed.Data.ID, parsed.Data.AccessToken)

	return parsed.Data.AccessToken, err
}

func parseAuthResp(body []byte) (authresp, error) {
	p := authresp{}
	err := json.Unmarshal(body, &p)
	if (err != nil) {
		panic(err.Error())
	}

	return p, err
}

func genURL(base string, urlpath string) (string) {
	u, err := url.Parse(base)
	if (err != nil) {
		panic(err.Error())
	}

	u.Path = path.Join(u.Path, urlpath)
	return u.String()
}
	
func getDefaultHeaders() map[string]string {
	return map[string]string {
		"Content-Type": "application/json",
		"Accept": "*/*",
		"Accept-Encoding": "gzip,deflate,br",
		"Connection": "keep-alive",
	}
}

// NewWmClient to construct waste management client
func NewWmClient(conf Config, restClient RestClient) *WMClient {
	return &WMClient{
		conf: conf,
		restClient: restClient,
	}
}