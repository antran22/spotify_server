package routes

import (
	"log"
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, _ *http.Request) {
	renderer := GetTemplateRenderer()
	err := renderer.Render(w, "index.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
