package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/home", nil)

	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()

	HomePageHandler(responseRecorder, request)

	cupaloy.SnapshotT(t, responseRecorder.Body.String())
}
