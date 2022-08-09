package csapi // import "go.bnck.me/csapi"

import (
	"go.bnck.me/csapi/v2/apis/stat"
	"go.bnck.me/csapi/v2/internal/transport"
)

// Client represents the main client
type Client struct {
	transport *transport.Client

	stat stat.Client
}

// Stat returns a client related to stats
func (c *Client) Stat() stat.Client {
	if c.stat == nil {
		c.stat = stat.New(c.transport)
	}

	return c.stat
}
