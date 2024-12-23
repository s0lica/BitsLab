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
	query = "CREATE TABLE IF NOT EXISTS Users(uid int AUTO_INCREMENT PRIMARY KEY, Name VARCHAR(500), Email VARCHAR(500), Username VARCHAR(500) ,Password VARCHAR(500))"
	create, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(create)
	fmt.Println(1)
	query = "CREATE TABLE IF NOT EXISTS UserSessions(session_id VARCHAR(255) PRIMARY KEY, uid int NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, expires_at TIMESTAMP)"
	create, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
}
