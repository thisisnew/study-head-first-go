package main

import (
	"log"
	"net/http"
)

func write(w http.ResponseWriter, message string) {
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(w http.ResponseWriter, r *http.Request) {
	write(w, "Hello, web!")
}

func frenchHandler(w http.ResponseWriter, r *http.Request) {
	write(w, "Salute, web!")
}

func hindiHandler(w http.ResponseWriter, r *http.Request) {
	write(w, "Namaste, web!")
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salute", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
