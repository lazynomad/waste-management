package restclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// HTTPClient interface
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// RestClient to submit requests
type RestClient struct {
	client HTTPClient
}

// Get sends an HTTP GET request to the endpoint.
// Returns:
// 		Response code
//		Response body
//		Error code, when failed
func (rest *RestClient) Get(url string, body []byte, headers map[string]string) (int, []byte, error) {
	req, _ := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(body))
	return rest.send(req, headers)
}

// Post sends an HTTP POST request to the endpoint.
// Returns:
// 		Response code
//		Response body
//		Error code, when failed
func (rest *RestClient) Post(url string, body []byte, headers map[string]string) (int, []byte, error) {
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	return rest.send(req, headers)
}

// NewRestClient to construct the rest client
func NewRestClient(httpClient HTTPClient) RestClient {
	return RestClient{client: httpClient}
}

func (rest *RestClient) send(req *http.Request, headers map[string]string) (int, []byte, error) {
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
