package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// HTTPClient interface that can be mocked
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// RestClient to submit requests
type RestClient struct {
	client HTTPClient
}

// Get call submits a request to the endpoint and returns the response
func (rest *RestClient) Get(url string, body []byte, headers map[string]string) (int, []byte, error) {
	req, _ := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(body))
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := rest.client.Do(req)
	if (err != nil) {
		panic(err.Error())
	}
	defer res.Body.Close()

	resBody, _ := ioutil.ReadAll(res.Body)
	return res.StatusCode, resBody, err
}

// NewRestClient to construct rest client
func NewRestClient(httpClient HTTPClient) *RestClient {
	return &RestClient{client: httpClient}
}
