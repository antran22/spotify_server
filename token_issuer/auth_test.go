package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthRedirectURL(t *testing.T) {
	redirectURL := GetAuthRedirectURL()
	assert.NotEmpty(t, redirectURL)
	assert.Contains(t, redirectURL, "https://accounts.spotify.com/authorize")
}

func TestAuthPageCallbackHandler(t *testing.T) {
	t.Run("test without state", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/callback", nil)

		assert.NoError(t, err)

		responseRecorder := httptest.NewRecorder()

		AuthCallBackHandler(responseRecorder, request)

		assert.NotEmpty(t, responseRecorder.Body.String())
		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	})

}
