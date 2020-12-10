package csapi_test

import (
	"github.com/jpbede/csapi-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Stat(t *testing.T) {
	c, err := csapi.New("abc", "abc-123")
	assert.NoError(t, err)

	statAPI := c.Stat()
	if statAPI == nil {
		t.Error("Failed to get 'Stat' endpoint")
	}
}
