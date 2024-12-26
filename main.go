/// Routes

package main

import (
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/s0lica/BitsLab/dbbuilder"
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

func main() {
	//STYLESHEETS
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("./stylesheets"))))
	//ROUTING NO LOGIN
	http.HandleFunc("/login", Login)
	http.HandleFunc("/sign_up", Sign_up)
	http.HandleFunc("/Login_user", auth.Login_user)
	http.HandleFunc("/Sign_up_user", auth.Sign_up_user)
	http.HandleFunc("/", Index)
	http.HandleFunc("/probleme", Problem)
	http.HandleFunc("/logouthandle", auth.LogoutHandler)
	//DYNAMIC ROUTING
	http.HandleFunc("/problems/{id}", ProblemHandler)
	//ROUTING AUTHREQUIRED
	http.HandleFunc("/user", auth.AuthRequired(Userpage))
	//ROUTING ADMINREQUIRED
	http.HandleFunc("/admin/create_problem", auth.AdminRequired(Create_problem))
	//TABLES BUILDER
	dbbuilder.Build_databases()
	//sv start
	http.ListenAndServe(":8000", nil)
}
