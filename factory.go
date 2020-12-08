package csapi

import (
	"github.com/jpbede/csapi-go/internal/transport"
	"net/http"
)

type APIEndpoint string

const (
	APIEndpointProduction APIEndpoint = "https://monitor.certified-senders.org/api/v1"
	APIEndpointBeta       APIEndpoint = "https://monitor-beta.certified-senders.org/api/v1"
)

// New creates a new Client with APIUrl and APIKey
func New() *Client {
	return NewWithOptions(WithAPIEndpoint(APIEndpointProduction))
}

// NewWithClient creates a new Client with a given http.Client
func NewWithClient(httpClient *http.Client) *Client {
	return NewWithOptions(WithAPIEndpoint(APIEndpointProduction), WithHTTPClient(httpClient))
}

// New creates a new Client with given options
func NewWithOptions(options ...ClientOption) *Client {
	c := &Client{}

	// always create a base transport it can be overwritten with options
	c.transport = transport.NewClient(string(APIEndpointProduction), nil)

	// run given options
	for _, option := range options {
		option(c)
	}

	return c
}
