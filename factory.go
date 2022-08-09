package csapi

import (
	"errors"
	"go.bnck.me/csapi/internal/transport"
	"net/http"
)

// APIEndpoint represents a api endpoint
type APIEndpoint string

const (
	// APIEndpointProduction Production API endpoint
	APIEndpointProduction APIEndpoint = "https://monitor.certified-senders.org/api/v1"
	// APIEndpointBeta Beta API endpoint
	APIEndpointBeta APIEndpoint = "https://monitor-beta.certified-senders.org/api/v1"
)

// New creates a new Client with APIUrl and APIKey
func New(apiKey string) (*Client, error) {
	return NewWithOptions(WithAPIEndpoint(APIEndpointProduction), WithAPIKey(apiKey))
}

// NewWithClient creates a new Client with a given http.Client
func NewWithClient(httpClient *http.Client, apiKey string) (*Client, error) {
	return NewWithOptions(WithAPIEndpoint(APIEndpointProduction), WithHTTPClient(httpClient), WithAPIKey(apiKey))
}

// NewWithOptions creates a new Client with given options
func NewWithOptions(options ...ClientOption) (*Client, error) {
	c := &Client{}

	// always create a base transport it can be overwritten with options
	c.transport = transport.NewClient(string(APIEndpointProduction), nil)

	// run given options
	for _, option := range options {
		option(c)
	}

	// check if there are credentials and then login
	if c.transport.APIKey == "" {
		return nil, errors.New("no api key supplied")
	}

	return c, nil
}
