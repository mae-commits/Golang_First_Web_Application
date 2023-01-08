package main

import (
	"fmt"
	"handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.ViewHandler)
	http.HandleFunc("/create", handler.CreateHandler)
	http.HandleFunc("/delete", handler.DeleteHandler)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	fmt.Println("Server Start Up........")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
