package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://google.com",
		"https://rpsconsulting.in",
		"https://stackoverflow.com",
	}
	//create the channel
	c := make(chan string)
	fmt.Println("Main Started")
	for _, link := range links {
		go ReadLinks(link, c)
	}
	//receiver
	for _, link := range links {
		time.Sleep(5 * time.Second)
		fmt.Println(link, <-c)
	}
	//fmt.Println(<-c)
	fmt.Println("Main Done")
}

func ReadLinks(link string, c chan string) {
	response, err := http.Get(link)
	if err != nil {
		fmt.Println("The site is down", link)
		//write the message to channel
		c <- "The site is down"
	} else {
		fmt.Println("The site is up", link)
		go ReadUrl(response)
		//write message to channel
		c <- "The site is up"
	}
}

func ReadUrl(response *http.Response) {

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Could not read contents", err)
	}
	fmt.Printf("The Content length %d", len(body))
}
