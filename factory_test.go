package csapi_test

import (
  "github.com/jpbede/csapi-go"
  "github.com/stretchr/testify/assert"
  "net/http"
  "testing"
)

func TestNew(t *testing.T) {
  c := csapi.New()
  assert.NotNil(t, c)
}

func TestNewWithClient(t *testing.T) {
  c := csapi.NewWithClient(&http.Client{})
  assert.NotNil(t, c)
}

func TestNewWithOptions(t *testing.T) {
  c := csapi.NewWithOptions(csapi.WithAPIEndpoint(""))
  assert.NotNil(t, c)
}
