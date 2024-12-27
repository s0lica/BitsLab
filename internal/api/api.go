package api

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/s0lica/BitsLab/internal/auth"
	"github.com/s0lica/BitsLab/internal/db"
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

func panicerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Create_problemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form["problem_name"]
	strtime_limit := r.Form["time_limit"]
	strmemory_limit := r.Form["memory_limit"]
	strsource_size := r.Form["source_size"]
	strdifficulty := r.Form["difficulty"]
	visible := r.FormValue("checkbox-visible") == "true"
	visible_tests := r.FormValue("checkbox-visible-tests") == "true"
	console_input := r.FormValue("checkbox-console") == "true"
	task_description := r.Form["task_description"]
	time_limit, err := strconv.ParseFloat(strtime_limit[0], 64)
	panicerr(err)
	memory_limit, err := strconv.Atoi(strmemory_limit[0])
	panicerr(err)
	source_size, err := strconv.ParseFloat(strsource_size[0], 64)
	panicerr(err)
	difficulty, err := strconv.Atoi(strdifficulty[0])
	panicerr(err)
	db.InitDB()
	_, err = db.DB.Query(`INSERT INTO Problems 
	(name,
	time_limit
	,memory_limit
	,source_size
	,console_input
	,visible
	,visible_tests
	,task_description
	,difficulty) VALUES (?,?,?,?,?,?,?,?,?)`,
		(name[0]),
		(time_limit),
		(memory_limit),
		(source_size),
		(console_input),
		(visible),
		(visible_tests),
		(task_description[0]),
		(difficulty))
	db.CloseDB()
	panicerr(err)
	session, _ := auth.Store.Get(r, "bitslab-session")
	username := session.Values["username"]
	data := map[string]interface{}{
		"username": username,
	}
	var tmpl = template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, data)
}
