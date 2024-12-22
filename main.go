package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("templates/log-in.html"))
	tmpl.Execute(w, nil)
}

func sign_up(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("templates/sign_up.html"))
	tmpl.Execute(w, nil)
}

func problem(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("templates/probleme.html"))
	tmpl.Execute(w, nil)
}

func sign_up_user(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var username = r.Form["username"]
	var email = r.Form["email"]
	var name = r.Form["name"]
	var pass = r.Form["password"]
	var repass = r.Form["repassword"]
	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(username)
	fmt.Println(pass)
	fmt.Println(repass)
	var tmpl = template.Must(template.ParseFiles("templates/sign_up.html"))
	tmpl.Execute(w, nil)
}

func login_user(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var username = r.Form["username"]
	var pass = r.Form["password"]
	fmt.Println(username)
	fmt.Println(pass)
	var tmpl = template.Must(template.ParseFiles("templates/log-in.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))
	http.HandleFunc("/", login)
	http.HandleFunc("/sign_up", sign_up)
	http.HandleFunc("/login_user", login_user)
	http.HandleFunc("/sign_up_user", sign_up_user)
	http.HandleFunc("/index", index)
	http.HandleFunc("/probleme", problem)
	http.ListenAndServe(":8000", nil)
}
