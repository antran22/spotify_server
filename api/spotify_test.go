package api

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSpotifyToken(t *testing.T) {

	input := []byte(`{
	"access_token": "123", 
	"token_type": "Bearer",
	"expiry": "2022-05-19T20:32:09.886172+07:00",
	"refresh_token": "123"
}`)

	token := parseSpotifyToken(bytes.NewReader(input))

	assert.NotEmpty(t, token)

	assert.Equal(t, "123", token.AccessToken)
	assert.Equal(t, "Bearer", token.TokenType)
	assert.NotEmpty(t, token.Expiry)
	assert.Equal(t, "123", token.RefreshToken)
}

func TestGetSpotifyToken(t *testing.T) {
	setEnvErr := os.Setenv("SPOTIFY_TOKEN_FILE", "../token.json")

	assert.NoError(t, setEnvErr)

	token := GetSpotifyToken()

	assert.NotEmpty(t, token)

	t.Cleanup(func() {
		unsetEnvError := os.Unsetenv("SPOTIFY_TOKEN_FILE")
		if unsetEnvError != nil {
			panic(unsetEnvError)
		}
	})
}
