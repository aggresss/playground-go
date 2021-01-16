package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoStuffWithTestServer(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "/some/path")
		// Send response to be tested
		rw.Write([]byte(`OK`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	api := API{server.Client(), server.URL}
	body, err := api.DoStuff()

	assert.Nil(t, err)
	assert.Equal(t, []byte("OK"), body)

}
