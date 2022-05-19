package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerIntegration(t *testing.T) {
	guardIntegrationTest(t)

	t.Log("opening test server")
	testServer := getTestServer()

	t.Cleanup(func() {
		t.Log("closing test server")
		testServer.Close()
	})

	t.Run("test /ping route", func(t *testing.T) {
		t.Parallel()
		resp, body := makeTestRequest(t, testServer, http.MethodGet, "/ping", nil)

		assert.Equal(t, resp.StatusCode, http.StatusOK)
		assert.Equal(t, "pong", body, "response not matching")
	})

	t.Run("test / route", func(t *testing.T) {
		t.Parallel()
		resp, _ := makeTestRequest(t, testServer, http.MethodGet, "/", nil)
		assert.Equal(t, resp.StatusCode, http.StatusOK, "status is not ok")
	})
}
