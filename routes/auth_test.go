package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthRedirectUrl(t *testing.T) {
	assert.NotEmpty(t, getAuthRedirectUrl())
}

func TestAuthPageHandler(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/auth", nil)

	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()

	AuthPageHandler(responseRecorder, request)
	response := responseRecorder.Result()
	assert.Equal(t, http.StatusFound, response.StatusCode)

	location, err := response.Location()
	assert.NoError(t, err)
	assert.Equal(t, getAuthRedirectUrl(), location.String())
}

func TestAuthPageCallbackHandler(t *testing.T) {
	t.Run("test without state", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/auth/callback", nil)

		assert.NoError(t, err)

		responseRecorder := httptest.NewRecorder()

		AuthCallBackHandler(responseRecorder, request)

		assert.NotEmpty(t, responseRecorder.Body.String())
		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	})

}
