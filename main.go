package main

import (
	handlers "Golang_First_Web_Application/handlers"
	"net/http"
)

func main() {
	// Define main page.
	http.HandleFunc("/", handlers.MainHandler)
	// http.Handle("/", http.FileServer(http.FS(handlers.Pages)))
	http.HandleFunc("/nextPage", handlers.NextHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./handlers/Pages/"))))
	http.ListenAndServe(":8080", nil)
}
