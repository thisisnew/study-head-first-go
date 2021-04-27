package main

import (
	"html/template"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(w, nil)
	check(err)
}

func newHanlder(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(w, nil)
	check(err)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	signature := r.FormValue("signature")
	_, err := w.Write([]byte(signature))
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHanlder)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe(":8000", nil)
	log.Fatal(err)
}
