package main

import (
	"fmt"
	"github.com/goadvanced/globalinsurance/stores"
)

func main() {
	/*
	   	vehicle:=models.Vehicle{
	   		"TN-02-483568","Honda",
	   		&models.Date{12,12,2020},"A23452","K42545","Diesel",
	   		"Blue",
	   	}

	     fmt.Printf("The Vehicle Details are %+v", vehicle)
	*/

	db := stores.ConnectionHelper()

	if db != nil {
		fmt.Println("Connected")
	}

}
