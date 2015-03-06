package bitly

import (
	"net/http"
	"net/url"
)

const (
	baseDefaultURL = "https://api-ssl.bitly.com"
	versionDefault = "/v3"
)

// A client manages communication with the bitly API.
type Client struct {
	// HTTP Client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// Services
	LinkMetric *LinkMetricService
}

// NewClient returns a new Bitly API client. If a nil httpClient is
// provided, http.DefaultClient will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the goauth2 library).
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(baseDefaultURL + versionDefault)

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.LinkMetric = &LinkMetricService{client: c}

	return c
}
