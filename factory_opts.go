package csapi

import (
	"github.com/jpbede/csapi-go/internal/transport"
	"github.com/rs/zerolog"
	"net/http"
)

type ClientOption func(*Client)

func WithAPIEndpoint(api APIEndpoint) ClientOption {
	return func(c *Client) {
		c.transport.BaseURL = string(api)
		c.transport = transport.NewClient(string(api), nil)
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.transport = transport.NewClient(c.transport.BaseURL, httpClient)
	}
}

func WithLogger(logger *zerolog.Logger) ClientOption {
	return func(c *Client) {
		c.transport.SetLogger(logger)
	}
}
