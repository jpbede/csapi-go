package csapi

import (
	"github.com/jpbede/csapi-go/apis/stat"
	"github.com/jpbede/csapi-go/internal/transport"
)

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
