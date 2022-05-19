package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func getMainRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", HomePageHandler)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			log.Println(err)
		}
	})
	return r
}

func main() {
	server := getMainRouter()
	log.Println("Listening at :3000")
	log.Fatalln(http.ListenAndServe(":3000", server))
}
