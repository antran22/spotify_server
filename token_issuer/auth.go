package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	spotifyAuth "github.com/zmb3/spotify/v2/auth"
)

const redirectURL = "http://localhost:3000/callback"
const state = "123"

func getSpotifyAuth() *spotifyAuth.Authenticator {
	clientId := os.Getenv("SPOTIFY_ID")
	if clientId == "" {
		log.Fatalln("Must Specify a SPOTIFY_ID environment variable")
	}

	clientSecret := os.Getenv("SPOTIFY_SECRET")
	if clientSecret == "" {
		log.Fatalln("Must Specify a SPOTIFY_SECRET environment variable")
	}

	return spotifyAuth.New(
		spotifyAuth.WithClientID(clientId),
		spotifyAuth.WithClientSecret(clientSecret),
		spotifyAuth.WithRedirectURL(redirectURL),
		spotifyAuth.WithScopes(spotifyAuth.ScopeStreaming,
			spotifyAuth.ScopeUserReadPlaybackState,
			spotifyAuth.ScopeUserReadPrivate,
		),
	)
}

var auth = getSpotifyAuth()

func GetAuthRedirectURL() string {
	url := auth.AuthURL(state)
	return url
}

func AuthCallBackHandler(w http.ResponseWriter, r *http.Request) {
	// use the same state string here that you used to generate the URL
	token, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "couldn't get token", http.StatusBadRequest)
		return
	}

	jsonToken, err := json.Marshal(token)
	if err != nil {
		http.Error(w, "couldn't marshal json", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonToken)
	if err != nil {
		log.Println(err)
		return
	}
}
