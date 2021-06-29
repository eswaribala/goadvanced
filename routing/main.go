package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "7070"
)

func home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Launching Home Page")
}
func product(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Launching Product Page")
}

func aboutUs(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Launching About Page")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/product", product)
	http.HandleFunc("/aboutus", aboutUs)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("Error in connection establishment", err)
	}
	return
}
