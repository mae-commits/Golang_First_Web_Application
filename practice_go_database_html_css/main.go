package main

import (
	"log"
	"net/http"

	handler "handler"
)

func main() {
	http.HandleFunc("/", handler.MainHandler)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
