package dbbuilder

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Build_databases() {
	db, err := sql.Open("mysql", "root:#David2007vasiliu@tcp(127.0.0.1)/BitsLab")
	if err != nil {
		panic(err)
	}
	var query string
	query = `CREATE TABLE IF NOT EXISTS Users
	(uid int AUTO_INCREMENT PRIMARY KEY, 
	Name VARCHAR(500), 
	Email VARCHAR(500), 
	Username VARCHAR(500), 
	Password VARCHAR(500),
	IsAdmin bool DEFAULT false)`
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(1)
	query = "CREATE TABLE IF NOT EXISTS UserSessions(session_id VARCHAR(255) PRIMARY KEY, uid int NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, expires_at TIMESTAMP)"
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	query = `CREATE TABLE IF NOT EXISTS Problems 
	(ID bigint AUTO_INCREMENT PRIMARY KEY, 
	created_at timestamp NOT NULL DEFAULT NOW(),
	name text NOT NULL,
	test_name text,
	time_limit double NOT NULL,
	memory_limit int NOT NULL,
	source_size int, 
	console_input bool,
	visible bool,
	visible_tests bool,
	task_description text,
	difficulty int NOT NULL)`
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	query = `CREATE TABLE IF NOT EXISTS Submissions 
	(ID bigint AUTO_INCREMENT PRIMARY KEY,
	code text NOT NULL,
	created_at timestamp NOT NULL DEFAULT NOW(),
	user_id int REFERENCES Users(uid),
	problem_id bigint REFERENCES Problems(ID),
	compile_error bool,
	compile_message text,
	compile_duration double,
	score int NOT NULL DEFAULT 0,
	max_time double,
	max_memory int)
	`
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	query = `CREATE TABLE IF NOT EXISTS TestCases
	(ID bigint AUTO_INCREMENT PRIMARY KEY,
	create_at timestamp NOT NULL DEFAULT NOW(),
	problem_id bigint REFERENCES Problems(ID),
	input longtext NOT NULL,
	expected_output longtext NOT NULL,
	visible bool,
	inproblemid int)`
	_, err = db.Exec(query)
	query = `CREATE TABLE IF NOT EXISTS SubmissionTestResults
	(ID bigint AUTO_INCREMENT PRIMARY KEY,
	submission_id bigint REFERENCES Submissions(ID),
	test_id bigint REFERENCES TestCases(ID),
	score int NOT NULL DEFAULT 0,
	time double,
	memory int,
	output text,
	compile_error bool,
	compile_message text,
	compile_duration double)`
	_, err = db.Exec(query)
}
