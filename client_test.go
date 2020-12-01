package csapi_test

import (
  "bytes"
  "context"
  "github.com/jpbede/csapi-go"
  "github.com/stretchr/testify/assert"
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestClient_Login(t *testing.T) {
  server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
    assert.Equal(t, req.URL.String(), "/login")

    body, readErr := ioutil.ReadAll(req.Body)
    assert.NoError(t, readErr)
    assert.Equal(t, `{"name":"test123","key":"test123"}`, string(bytes.Trim(body, "\n")))
  }))

  c := csapi.NewWithOptions(csapi.WithAPIEndpoint(csapi.APIEndpoint(server.URL)), csapi.WithHTTPClient(server.Client()))
  c.Login(context.Background(), "test123", "test123")
}

func TestClient_Stat(t *testing.T) {
  c := csapi.New()
  statAPI := c.Stat()
  if statAPI == nil {
    t.Error("Failed to get 'Stat' endpoint")
  }
}
