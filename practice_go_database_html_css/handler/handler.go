package handler

import (
	"log"
	"net/http"
	"text/template"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("topPage.html")
	if err != nil {
		log.Fatal(err)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {

}
