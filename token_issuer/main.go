package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func getMainRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/callback", AuthCallBackHandler)
	return r
}

func main() {
	server := getMainRouter()
	log.Printf("Open your browser and go to: %s\n", GetAuthRedirectURL())
	log.Println("Listening at :3000")
	log.Fatalln(http.ListenAndServe(":3000", server))
}
