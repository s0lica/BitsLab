package main

import (
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("templates/log-in.html"))
	tmpl.Execute(w, nil)
}

func sign_up(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("templates/sign_up.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc("/sign_up", sign_up)
	http.ListenAndServe(":8000", nil)
}
