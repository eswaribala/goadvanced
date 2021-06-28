package main

import (
	"fmt"
	"net/http"
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
	fmt.Println(<-c)
	fmt.Println("Main Done")
}

func ReadLinks(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("The site is down", link)
		//write the message to channel
		c <- "The site is down"
	} else {
		fmt.Println("The site is up", link)
		//write message to channel
		c <- "The site is up"
	}
}
