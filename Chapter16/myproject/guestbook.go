package main

import (
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	placeholder := []byte("signature list goes here")
	_, err := w.Write(placeholder)
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe(":8000", nil)
	log.Fatal(err)
}
