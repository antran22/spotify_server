package routes

import (
	"encoding/json"
	"log"
	"net/http"

	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

const redirectURL = "http://localhost:3000/auth/callback"
const state = "123"

var auth = spotifyauth.New(
	spotifyauth.WithClientID("e6867a62802c4a298afb60e97c122a89"),
	spotifyauth.WithClientSecret("b72855e2cd444c72a34f93efebfbaee7"),
	spotifyauth.WithRedirectURL(redirectURL),
	spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate),
)

func getAuthRedirectUrl() string {
	// get the user to this URL - how you do that is up to you.
	// you should specify a unique state string to identify the session
	url := auth.AuthURL(state)
	log.Println(url)
	return url
}

func AuthPageHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, getAuthRedirectUrl(), http.StatusFound)
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
