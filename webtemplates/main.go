package main

import (
	"fmt"
	"github.com/goadvanced/webtemplates/models"
	"log"
	"net/http"
	"text/template"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "7070"
)

func home(w http.ResponseWriter, r *http.Request) {
	vehicle := models.Vehicle{RegNo: "TN-32-826486", Brand: "Honda", ModelName: "CIVIC"}
	parsedTemplate, _ := template.ParseFiles("templates/home.html")
	err := parsedTemplate.Execute(w, vehicle)
	if err != nil {
		fmt.Println("Error occurred while executing the templateor writing its output : ", err)
		return
	}
}

func main() {
	fmt.Println("Connection Ready at http://" + CONN_HOST + ":" + CONN_PORT)
	http.HandleFunc("/", home)

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)

	if err != nil {
		log.Fatal("Error in connection establishment", err)
	}
	return
}
