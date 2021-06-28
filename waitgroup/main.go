package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var wg1 sync.WaitGroup

func main() {

	links := []string{
		"https://google.com",
		"https://rpsconsulting.in",
		"https://stackoverflow.com",
	}

	fmt.Println("Main Started")
	wg.Add(len(links)) //counter value is 3
	for _, link := range links {
		go ReadLinks(link)
	}

	wg.Wait()
	fmt.Println("\nMain Done")
}

func ReadLinks(link string) {
	defer wg.Done()
	wg1.Add(1)
	response, err := http.Get(link)

	if err != nil {
		fmt.Println("The site is down", link)

	}
	//defer response.Body.Close()
	fmt.Println("The site is up", link)
	go ReadUrl(response)
	wg1.Wait()
	fmt.Printf("\nThe Response is received from %s", link)

	//wg.Done()
}

func ReadUrl(response *http.Response) {
	defer wg1.Done()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Could not read contents", err)
	}
	fmt.Printf("The Content length \n%d\n", len(body))
}
