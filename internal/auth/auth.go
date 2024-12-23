package auth

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func hashPassword(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func add_user(name string, email string, username string, password string, repassword string) bool {
	db, err := sql.Open("mysql", "root:#David2007vasiliu@tcp(127.0.0.1)/BitsLab")
	fmt.Println(1)
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

func Sign_up_user(w http.ResponseWriter, r *http.Request) {
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

func Login_user(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var username = r.Form["username"]
	var pass = r.Form["password"]
	if check_user(username[0], pass[0]) {
		session, _ := store.Get(r, "bitslab-session")
		session.Values["authenticated"] = true
		session.Values["username"] = username
		err := session.Save(r, w)
		if err != nil {
			panic(err)
		}
		var tmpl = template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	} else {
		var tmpl = template.Must(template.ParseFiles("templates/probleme.html"))
		tmpl.Execute(w, nil)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "bitslab-session")
	for k := range session.Values {
		delete(session.Values, k)
	}
	session.Save(r, w)
	var tmpl = template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}
