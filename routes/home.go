package routes

import (
	"log"
	"net/http"

	"spotify_server/template"
)

func HomePageHandler(w http.ResponseWriter, _ *http.Request) {
	renderer := template.GetTemplateRenderer()
	err := renderer.Render(w, "index.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
