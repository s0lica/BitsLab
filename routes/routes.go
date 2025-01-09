package routes

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/s0lica/BitsLab/internal/auth"
	"github.com/s0lica/BitsLab/internal/db"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
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
	fmt.Println(session.Values["autenticated"])
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
	db.InitDB()
	query := fmt.Sprintf(`SELECT name, 
						 time_limit, 
						 memory_limit, 
						 console_input, 
						 task_description, 
						 difficulty 
						 FROM Problems WHERE ID = '%s'`, (problemid))
	var name string
	var time_limit float64
	var memory_limit int
	var console_input bool
	var task_description string
	var difficulty int
	err := db.DB.QueryRow(query).Scan(&name, &time_limit, &memory_limit, &console_input, &task_description, &difficulty)
	if err != nil {
		http.Error(w, "Problem does not exist", http.StatusForbidden)
		return
	}
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Strikethrough,
			extension.Table,
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	var buf bytes.Buffer
	err = md.Convert([]byte(task_description), &buf)
	fmt.Println(buf.String())
	if err != nil {
		http.Error(w, "Failed to render md", http.StatusForbidden)
		return
	}
	db.CloseDB()
	data := map[string]interface{}{
		"username":         username,
		"problemid":        problemid,
		"name":             name,
		"time_limit":       time_limit,
		"memory_limit":     memory_limit,
		"console_input":    console_input,
		"task_description": template.HTML(buf.String()),
		"difficulty":       difficulty,
	}
	var tmpl = template.Must(template.ParseFiles("templates/problems/problem_template.html"))
	tmpl.Execute(w, data)
}

func Edit_problem(w http.ResponseWriter, r *http.Request) {
	session, _ := auth.Store.Get(r, "bitslab-session")
	username := session.Values["username"]
	problemid := r.PathValue("id")
	db.InitDB()
	query := fmt.Sprintf(`SELECT name,
						 time_limit,
						 memory_limit,
						 source_size,
						 console_input,
						 visible,
						 visible_tests,
						 task_description,
						 difficulty
						 FROM Problems WHERE ID = '%s'`, (problemid))
	var name string
	var time_limit float64
	var memory_limit int
	var source_size float64
	var console_input bool
	var visible bool
	var visible_tests bool
	var task_description string
	var difficulty int
	err := db.DB.QueryRow(query).Scan(&name, &time_limit, &memory_limit, &source_size, &console_input, &visible, &visible_tests, &task_description, &difficulty)
	if err != nil {
		http.Error(w, "Problem does not exist", http.StatusForbidden)
	}
	db.CloseDB()
	data := map[string]interface{}{
		"username":         username,
		"problemid":        problemid,
		"name":             name,
		"time_limit":       time_limit,
		"memory_limit":     memory_limit,
		"console_input":    console_input,
		"task_description": task_description,
		"difficulty":       difficulty,
		"visible":          visible,
		"visible_tests":    visible_tests,
		"source_size":      source_size,
	}
	var tmpl = template.Must(template.ParseFiles("templates/problems/edit_problem.html"))
	tmpl.Execute(w, data)
}
