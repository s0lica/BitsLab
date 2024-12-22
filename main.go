package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func add_user(name string, email string, username string, password string, repassword string) bool {
	db, err := sql.Open("mysql", "root:#David2007vasiliu@tcp(127.0.0.1)/BitsLab")
	if password != repassword {
		return false
	}
	if err != nil {
		panic(err)
	}
	add, err := db.Query("INSERT INTO Users(Name,Email,Username,Password) VALUES (?,?,?,?)", (name), (email), (username), (password))
	if err != nil {
		panic(err)
	}
	fmt.Println(add)
	defer db.Close()
	return true
}

func check_user(username string, password string) bool {
	db, err := sql.Open("mysql", "root:#David2007vasiliu@tcp(127.0.0.1)/BitsLab")
	if err != nil {
		panic(err)
	}
	var exists bool
	var query string
	query = fmt.Sprintf("SELECT EXISTS (SELECT Name FROM Users WHERE Name='%s' AND Password='%s')", (username), (password))
	row := db.QueryRow(query).Scan(&exists)
	fmt.Println(row)
	defer db.Close()
	return exists
}

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
	if add_user(name[0], email[0], username[0], pass[0], repass[0]) {
		var tmpl = template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	} else {
		var tmpl = template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	}
}

func login_user(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var username = r.Form["username"]
	var pass = r.Form["password"]
	fmt.Println(username)
	fmt.Println(pass)
	if check_user(username[0], pass[0]) {
		var tmpl = template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	} else {
		var tmpl = template.Must(template.ParseFiles("templates/probleme.html"))
		tmpl.Execute(w, nil)
	}
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
