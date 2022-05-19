package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("templates/home.gohtml"))

func HomePageHandler(w http.ResponseWriter, _ *http.Request) {
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
