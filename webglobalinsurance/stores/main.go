package stores

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goadvanced/webglobalinsurance/domain"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func ConnectionHelper() {

	var err error
	dataSourceName := "root:vignesh@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	db.Exec("CREATE DATABASE amexdb")
	db.Exec("USE amexdb")

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&domain.PolicyHolder{}, &domain.Policy{}, &domain.Vehicle{})
}

//create PolicyHolder

func CreatePolicyHolder(policyHolder domain.PolicyHolder) domain.PolicyHolder {

	db.Create(policyHolder)
	return policyHolder
}

//Read all the PolicyHolders

func ReadAllPolicyHolders() ([]*domain.PolicyHolder, error) {
	var policyHolders []*domain.PolicyHolder
	db.Find(&policyHolders)
	return policyHolders, nil
}

//Update PolicyHolder

//Delete PolicyHolder
