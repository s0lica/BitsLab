package api

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/s0lica/BitsLab/internal/auth"
)

type problem struct {
	name             string
	test_name        string
	time_limit       float64
	memory_limit     int
	source_size      float64
	console_input    bool
	visible          bool
	visible_tests    bool
	task_description string
	difficulty       int
}

func Create_problemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form["problem_name"]
	time_limit := r.Form["time_limit"]
	memory_limit := r.Form["memory_limit"]
	source_size := r.Form["source_size"]
	difficulty := r.Form["difficulty"]
	visible := r.FormValue("checkbox-visible") == "true"
	visible_tests := r.FormValue("checkbox-visible-tests") == "true"
	task_description := r.Form["task_description"]
	fmt.Println(name)
	fmt.Println(memory_limit)
	fmt.Println(time_limit)
	fmt.Println(source_size)
	fmt.Println(difficulty)
	fmt.Println(visible)
	fmt.Println(visible_tests)
	fmt.Println(task_description)
	session, _ := auth.Store.Get(r, "bitslab-session")
	username := session.Values["username"]
	data := map[string]interface{}{
		"username": username,
	}
	var tmpl = template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, data)
}
