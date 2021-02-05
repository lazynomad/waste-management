package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"path"
	"github.com/lazynomad/waste-management/restclient"
)

// WMClient to communicate with waste management endpoint
type WMClient struct {
	conf config
	restClient restclient.RestClient
}

// GetAuthToken to get the JWT token from Waste-Management
// Returns:
//		JWT token
//		error, when failed
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
		errMsg := fmt.Sprintf("Unable to fetch auth token. Response code: %d. Body: %s", respCode, string(respBody))
		return "", errors.New(errMsg)
	}

	parsed, err := parseAuthResp(respBody)
	if (err != nil) {
		panic(err.Error())
	}

	// To-Do: Replace fmt with logger
	fmt.Printf("status: %d - id: %s - token:%s", parsed.StatusCode, parsed.Data.UserID, parsed.Data.AccessToken)
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
func NewWmClient(conf config, restClient restclient.RestClient) *WMClient {
	return &WMClient{
		conf: conf,
		restClient: restClient,
	}
}