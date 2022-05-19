package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
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
		resp, body := makeTestRequest(t, testServer, http.MethodGet, "/ping", nil)

		assert.Equal(t, resp.StatusCode, http.StatusOK)
		assert.Equal(t, "pong", body, "response not matching")
	})

	t.Run("test / route", func(t *testing.T) {
		resp, body := makeTestRequest(t, testServer, http.MethodGet, "/", nil)

		assert.Equal(t, resp.StatusCode, http.StatusOK, "status is not ok")

		cupaloy.SnapshotT(t, body)
	})

	t.Run("test /auth route", func(t *testing.T) {
		resp, _ := makeTestRequest(t, testServer, http.MethodGet, "/auth", nil)
		fmt.Println(resp)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		assert.Contains(t, resp.Request.URL.String(), "spotify")
	})
}
