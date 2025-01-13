package api

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
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
	row := db.DB.QueryRow("SELECT MAX(ID) FROM Problems")
	var last_id int
	row.Scan(&last_id)
	db.CloseDB()
	panicerr(err)
	fmt.Println(last_id)
	str := strconv.Itoa(last_id)
	http.Redirect(w, r, fmt.Sprintf("/admin/edit_problem/%s/create_test", (str)), http.StatusAccepted)
}

func Create_submissionHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
}

// / test creation from input / output format straight from form
func Create_testsimpleHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, (10 << 20))
	err := r.ParseMultipartForm(10 << 20)
	problem_id := r.PathValue("problem_id")
	fileinput, inputhandler, err := r.FormFile("infile")
	if err != nil {
		fmt.Println(err)
	}
	fileoutput, outputhandler, err := r.FormFile("outfile")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(inputhandler)
	fmt.Println(outputhandler)
	if fileinput != nil {
		defer fileinput.Close()
	}
	if fileoutput != nil {
		defer fileoutput.Close()
	}
	var stringinputcontent string
	var stringoutputcontent string
	if fileinput != nil {
		inputcontent, _ := io.ReadAll(fileinput)
		stringinputcontent = string(inputcontent)
	}
	if fileoutput != nil {
		outputcontent, _ := io.ReadAll(fileoutput)
		stringoutputcontent = string(outputcontent)
	}
	input := r.Form["testcase_input"]
	expected_output := r.Form["testcase_output"]
	var pid int
	if stringinputcontent != "" && input[0] == "" {
		input[0] = stringinputcontent
	}
	if stringoutputcontent != "" && expected_output[0] == "" {
		expected_output[0] = stringoutputcontent
	}
	pid, _ = strconv.Atoi(problem_id)
	db.InitDB()
	_, err = db.DB.Query(`INSERT INTO TestCases
	(problem_id,
	input,
	expected_output) VALUES (?,?,?)`,
		(pid),
		(input[0]),
		(expected_output[0]))
	if err != nil {
		fmt.Println(err)
	}
	db.CloseDB()
	rdr := fmt.Sprintf("/admin/edit_problem/%s/create_test", (problem_id))
	http.Redirect(w, r, rdr, http.StatusSeeOther)
}
