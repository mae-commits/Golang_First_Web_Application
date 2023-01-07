package main

import (
	"html/template"
	"log"
	"net/http"
)

// TemperatureDataElem is the struct which has the dataset of temperature
type TemperatureDataElem struct {
	Label string
	Data  []float64
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Read the file "mainPage.html"
	t, err := template.ParseFiles("./Pages/mainPage.html")
	if err != nil {
		log.Fatal(err)
	}
	// Output the content of "mainPage.html"
	if err := t.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func nextHandler(w http.ResponseWriter, r *http.Request) {
	var temperatureData []TemperatureDataElem
	temperatureData = append(temperatureData, TemperatureDataElem{
		Label: "沖縄県",
		Data:  []float64{5.2, 5.7, 13.9, 18.2, 21.4, 25.0, 26.4, 22.8, 17.5, 12.1, 7.6},
	})
	temperatureData = append(temperatureData, TemperatureDataElem{
		Label: "東京都",
		Data:  []float64{5.2, 5.7, 13.9, 18.2, 21.4, 25.0, 26.4, 22.8, 17.5, 12.1, 7.6},
	})
	t, err := template.ParseFiles("./Pages/nextPage.html")
	if err != nil {
		log.Fatal(err)
	}
	// Pass the data of "temperatureData"
	if err := t.Execute(w, temperatureData); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Define main page.
	http.HandleFunc("/", handler)
	http.HandleFunc("/nextPage", nextHandler)
	http.ListenAndServe(":8080", nil)
}
