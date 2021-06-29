package main

import (
	"encoding/json"
	"fmt"
	"github.com/goadvanced/webglobalinsurance/handlers"
	"github.com/goadvanced/webglobalinsurance/stores"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func main() {

	data, _ := ioutil.ReadFile("config.json")
	//fmt.Println(data)

	//define interface
	var result map[string]interface{}
	//convert
	json.Unmarshal([]byte(data), &result)
	host := result["host"].(string) + ":" + result["port"].(string)
	fmt.Println(host)

	stores.ConnectionHelper()
	router := mux.NewRouter()
	// Create
	router.HandleFunc("/policyholders", handlers.CreatePolicyHolderHandler).Methods("POST")
	// Read
	router.HandleFunc("/policyholders/{aadhaarCardNo}", handlers.GetPolicyHolderHandler).Methods("GET")
	// Read-all
	router.HandleFunc("/policyholders", handlers.ReadAllPolicyHoldersHandler).Methods("GET")
	// Update
	router.HandleFunc("/policyholders/{aadhaarCardNo}", handlers.UpdatePolicyHolderHandler).Methods("PUT")
	// Delete
	router.HandleFunc("/policyholders/{aadhaarCardNo}", handlers.DeletePolicyHolderHandler).Methods("DELETE")

	err := http.ListenAndServe(string(host), router)
	if err != nil {
		fmt.Println("Error in creating connection")
	}
	return

}
