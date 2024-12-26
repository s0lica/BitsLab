package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:#David2007vasiliu@tcp(127.0.0.1)/BitsLab")
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	log.Println("Connected to DB successfully!")
}

func CloseDB() {
	DB.Close()
}
