package csapi

import (
	"context"
	"errors"
	"github.com/jpbede/csapi-go/internal/transport"
	"net/http"
)

type APIEndpoint string

const (
	APIEndpointProduction APIEndpoint = "https://monitor.certified-senders.org/api/v1"
	APIEndpointBeta       APIEndpoint = "https://monitor-beta.certified-senders.org/api/v1"
)

// New creates a new Client with APIUrl and APIKey
func New(apiName, apiKey string) (*Client, error) {
	return NewWithOptions(WithAPIEndpoint(APIEndpointProduction), WithCredentials(apiName, apiKey))
}

// NewWithClient creates a new Client with a given http.Client
func NewWithClient(httpClient *http.Client, apiName, apiKey string) (*Client, error) {
	return NewWithOptions(WithAPIEndpoint(APIEndpointProduction), WithHTTPClient(httpClient), WithCredentials(apiName, apiKey))
}

// New creates a new Client with given options
func NewWithOptions(options ...ClientOption) (*Client, error) {
	c := &Client{}

	// always create a base transport it can be overwritten with options
	c.transport = transport.NewClient(string(APIEndpointProduction), nil)

	// run given options
	for _, option := range options {
		option(c)
	}

	// check if there are credentials and then login
	if !c.transport.HasCredentials() {
		return nil, errors.New("no api credentials supplied")
	}
	if err := c.transport.Login(context.Background()); err != nil {
		return nil, err
	}

	return c, nil
}
