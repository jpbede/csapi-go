package csapi_test

import (
	"github.com/stretchr/testify/assert"
	"go.bnck.me/csapi"
	"testing"
)

func TestClient_Stat(t *testing.T) {
	c, err := csapi.New("abc-123")
	assert.NoError(t, err)

	statAPI := c.Stat()
	if statAPI == nil {
		t.Error("Failed to get 'Stat' endpoint")
	}
}
