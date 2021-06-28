package main

import (
	"database/sql"
	"fmt"
)

func ConnectionHelper() *sql.DB {
	db, err := sql.Open("mysql",
		"root:vignesh@(127.0.0.1:3306)/amexglobalinsdb?parseTime=true")
	if err == nil {
		fmt.Println("Connection Error occurred", err)
	} else {
		fmt.Println("Connection Created")
	}

	return db
}