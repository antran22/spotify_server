package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2"
)

func GetSpotifyToken() *oauth2.Token {
	if spotifyToken != nil {
		return spotifyToken
	}
	spotifyToken = readSpotifyToken()
	return spotifyToken
}

var (
	spotifyToken *oauth2.Token
)

func readSpotifyToken() *oauth2.Token {
	files, err := ioutil.ReadDir("/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
	path := os.Getenv("SPOTIFY_TOKEN_FILE")
	if path == "" {
		path = "token.json"
	}

	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	token := parseSpotifyToken(jsonFile)
	return token
}

func parseSpotifyToken(r io.Reader) *oauth2.Token {
	byteValue, _ := ioutil.ReadAll(r)
	token := &oauth2.Token{}
	unmarshalError := json.Unmarshal(byteValue, token)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	return token
}
