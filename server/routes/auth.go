package routes

import (
	"log"
	"net/http"
)

func AuthPageHandler(w http.ResponseWriter, _ *http.Request) {
	renderer := GetTemplateRenderer()
	err := renderer.Render(w, "auth.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
