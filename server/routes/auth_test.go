package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/assert"
)

func TestAuthPageHandler(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/auth", nil)

	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()

	AuthPageHandler(responseRecorder, request)

	assert.NotEmpty(t, responseRecorder.Body.String())
	cupaloy.SnapshotT(t, responseRecorder.Body.String())
}
