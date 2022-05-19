package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"spotify_server/routes"
)

func getMainRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", routes.HomePageHandler)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			log.Println(err)
		}
	})
	r.Get("/auth", routes.AuthPageHandler)
	r.Get("/auth/callback", routes.AuthCallBackHandler)
	return r
}

func main() {
	server := getMainRouter()
	log.Println("Listening at :3000")
	log.Fatalln(http.ListenAndServe(":3000", server))
}
