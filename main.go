package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// Open database connection
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT * FROM table")
	if err != nil && err != sql.ErrNoRows {
		panic(err.Error())
	}

	if err == sql.ErrNoRows {
		log.Printf("not record")
	}

	log.Printf("rows%+v", rows)
}
