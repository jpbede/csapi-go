package csapi_test

import (
	"github.com/stretchr/testify/assert"
	"go.bnck.me/csapi"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	c, err := csapi.New("abc", "abc-123")
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestNewWithClient(t *testing.T) {
	c, err := csapi.NewWithClient(&http.Client{}, "abc", "abc-123")
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestNewWithOptions(t *testing.T) {
	c, err := csapi.NewWithOptions(csapi.WithAPIEndpoint(csapi.APIEndpointProduction), csapi.WithCredentials("abc", "abc-123"))
	assert.NoError(t, err)
	assert.NotNil(t, c)
}
