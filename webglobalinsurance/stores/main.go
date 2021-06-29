package stores

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goadvanced/webglobalinsurance/domain"
	"github.com/jinzhu/gorm"
)

func ConnectionHelper() *gorm.DB {

	connectionString := "root:vignesh@tcp(localhost:3306)/?parseTime=True"
	db, err := gorm.Open("mysql", connectionString)

	if err != nil {
		fmt.Println("Connection Not Established")
	}
	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	db.Exec("CREATE DATABASE amexdb")
	db.Exec("USE amexdb")

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&domain.PolicyHolder{}, &domain.Policy{}, domain.Vehicle{})

	return db
}
