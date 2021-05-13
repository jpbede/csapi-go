package stat

import (
	"go.bnck.me/csapi/internal/transport"
)

type client struct {
	transport *transport.Client
}

// New creates a new Client for the task Endpoint
func New(transport *transport.Client) Client {
	return &client{
		transport: transport,
	}
}
