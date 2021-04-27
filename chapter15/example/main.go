package main

import (
	"log"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	message := []byte("Hello, web!")
	_, err := w.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
