package main

import (
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
)

var templates = template.Must(template.ParseFiles("template.html"))

type page struct {
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := page{}
	templates.Execute(w, p)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public/")))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
