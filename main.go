// / Routes

package main

import (
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/s0lica/BitsLab/dbbuilder"
	"github.com/s0lica/BitsLab/internal/auth"
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
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("./stylesheets"))))
	http.HandleFunc("/login", login)
	http.HandleFunc("/sign_up", sign_up)
	http.HandleFunc("/Login_user", auth.Login_user)
	http.HandleFunc("/Sign_up_user", auth.Sign_up_user)
	http.HandleFunc("/", index)
	http.HandleFunc("/probleme", problem)
	dbbuilder.Build_databases()
	http.ListenAndServe(":8000", nil)
}
