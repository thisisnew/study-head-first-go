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

func d(w http.ResponseWriter, r *http.Request) {
	write(w, "z")
}

func e(w http.ResponseWriter, r *http.Request) {
	write(w, "x")
}

func f(w http.ResponseWriter, r *http.Request) {
	write(w, "y")
}

func main() {
	http.HandleFunc("/a", f)
	http.HandleFunc("/b", d)
	http.HandleFunc("/c", e)
	err := http.ListenAndServe(":4567", nil)
	log.Fatal(err)
}
