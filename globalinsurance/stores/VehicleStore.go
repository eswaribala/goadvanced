package stores

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectionHelper() *sql.DB {
	db, err := sql.Open("mysql",
		"root:vignesh@(127.0.0.1:3306)/amexglobalinsdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db

}
