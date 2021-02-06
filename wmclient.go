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
	headers := getHeaders(wm.conf.APIKeys.Auth, "")

	respCode, respBody, err := wm.restClient.Post(fullURL, jsonM, headers)
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
	fmt.Printf("Status: %d - User ID: %s - Token:%s", parsed.StatusCode, parsed.Data.UserID, parsed.Data.AccessToken)
	return parsed.Data.AccessToken, err
}

// GetAccountID to get the Account ID from Waste-Management
// Returns:
//		Account ID
//		error, when failed
func (wm *WMClient) GetAccountID(userID string, authToken string) (string, error) {
	suffixURL := fmt.Sprintf("authorize/user/%s/accounts?lang=en_US", userID)
	fullURL := genURL(wm.conf.BaseURL, suffixURL)
	headers := getHeaders(wm.conf.APIKeys.Account, authToken)
	
	respCode, respBody, err := wm.restClient.Get(fullURL, nil, headers)
	if (err != nil) {
		panic(err.Error())
	}
	
	if (respCode != 200) {
		fmt.Printf("Failed. Code: %d. Body: %s", respCode, string(respBody))
		errMsg := fmt.Sprintf("Unable to fetch Account ID. Response code: %d. Body: %s", respCode, string(respBody))
		return "", errors.New(errMsg)
	}

	parsed, err := parseAccountResp(respBody)
	if (err != nil) {
		panic(err.Error())
	}

	// To-Do: Replace fmt with logger
	fmt.Printf("Status: %d - Account ID: %s", parsed.StatusCode, parsed.Data.Accounts[0].ID)
	return parsed.Data.Accounts[0].ID, err
}

// GetServiceSchedules to get the scheduled service dates from waste maangement
// Returns:
//		Trash date ["02-12-2021"]
//		Recycling date ["02-12-2021"]
//		error, when failed
func (wm *WMClient) GetServiceSchedules(userID string, accountID string, authToken string) (string, string, error) {
	suffixURL := fmt.Sprintf("account/%s/service/1/pickupinfo?lang=en_US&userId=%s", accountID, userID)
	fullURL := genURL(wm.conf.BaseURL, suffixURL)
	headers := getHeaders(wm.conf.APIKeys.Service, authToken)
	
	respCode, respBody, err := wm.restClient.Get(fullURL, nil, headers)
	if (err != nil) {
		panic(err.Error())
	}
	
	if (respCode != 200) {
		fmt.Printf("Failed. Code: %d. Body: %s", respCode, string(respBody))
		errMsg := fmt.Sprintf("Unable to fetch Service schedules. Response code: %d. Body: %s", respCode, string(respBody))
		return "", "", errors.New(errMsg)
	}

	parsed, err := parseSchedulerResp(respBody)
	if (err != nil) {
		panic(err.Error())
	}

	// To-Do: Replace fmt with logger
	fmt.Printf("Next pickup date: %s. Message: %s", parsed.NextPickup.Date, parsed.NextPickup.Message)
	return parsed.NextPickup.Date, "", err
}

func parseAuthResp(body []byte) (authresp, error) {
	p := authresp{}
	err := json.Unmarshal(body, &p)
	if (err != nil) {
		panic(err.Error())
	}

	return p, err
}

func parseAccountResp(body []byte) (accountresp, error) {
	p := accountresp{}
	err := json.Unmarshal(body, &p)
	if (err != nil) {
		panic(err.Error())
	}

	return p, err
}

func parseSchedulerResp(body []byte) (scheduleresp, error) {
	p := scheduleresp{}
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
	
func getHeaders(apiKey string, authToken string) map[string]string {
	headers := map[string]string {
		"Content-Type": "application/json",
		"Accept": "*/*",
		"Accept-Encoding": "gzip,deflate,br",
		"Connection": "keep-alive",
		"apikey": apiKey,
	}
	if authToken != "" {
		headers["oktatoken"] = authToken
	}

	return headers
}

// NewWmClient to construct waste management client
func NewWmClient(conf config, restClient restclient.RestClient) WMClient {
	return WMClient{
		conf: conf,
		restClient: restClient,
	}
}