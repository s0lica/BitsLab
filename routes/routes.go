package routes

import (
	"html/template"
	"net/http"

	"github.com/s0lica/BitsLab/internal/auth"
)

func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := auth.Store.Get(r, "bitslab-session")
	username := session.Values["username"]
	data := map[string]interface{}{
		"username": username,
	}
	var tmpl = template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, data)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("templates/log-in.html"))
	tmpl.Execute(w, nil)
}

func Sign_up(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("templates/sign_up.html"))
	tmpl.Execute(w, nil)
}

func Problem(w http.ResponseWriter, r *http.Request) {
	session, _ := auth.Store.Get(r, "bitslab-session")
	username := session.Values["username"]
	data := map[string]interface{}{
		"username": username,
	}
	var tmpl = template.Must(template.ParseFiles("templates/probleme.html"))
	tmpl.Execute(w, data)
}

func Userpage(w http.ResponseWriter, r *http.Request) {
	session, _ := auth.Store.Get(r, "bitslab-session")
	username := session.Values["username"]
	data := map[string]interface{}{
		"username": username,
	}
	var tmpl = template.Must(template.ParseFiles("templates/userpage.html"))
	tmpl.Execute(w, data)
}

func Create_problem(w http.ResponseWriter, r *http.Request) {
	session, _ := auth.Store.Get(r, "bitslab-session")
	username := session.Values["username"]
	data := map[string]interface{}{
		"username": username,
	}
	var tmpl = template.Must(template.ParseFiles("templates/problems/create_problem.html"))
	tmpl.Execute(w, data)
}

func ProblemHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := auth.Store.Get(r, "bitslab-session")
	username := session.Values["username"]
	problemid := r.PathValue("id")
	data := map[string]interface{}{
		"username":  username,
		"problemid": problemid,
	}
	var tmpl = template.Must(template.ParseFiles("templates/problems/problem_template.html"))
	tmpl.Execute(w, data)
}
