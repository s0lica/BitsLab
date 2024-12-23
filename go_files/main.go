// / Routes and user parsing + user querying

package main

import (
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/s0lica/BitsLab/go_files/internal/auth/auth"
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

func main() {
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))
	http.HandleFunc("/login", login)
	http.HandleFunc("/sign_up", sign_up)
	http.HandleFunc("/login_user", login_user)
	http.HandleFunc("/sign_up_user", sign_up_user)
	http.HandleFunc("/", index)
	http.HandleFunc("/probleme", problem)
	http.ListenAndServe(":8000", nil)
}
