package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {

	links := []string{
		"https://google.com",
		"https://rpsconsulting.in",
		"https://stackoverflow.com",
	}

	fmt.Println("Main Started")
	wg.Add(len(links))
	for _, link := range links {
		go ReadLinks(link)
	}

	wg.Wait()
	fmt.Println("Main Done")
}

func ReadLinks(link string) {
	defer wg.Done()
	_, err := http.Get(link)
	//defer response.Body.Close()
	if err != nil {
		fmt.Println("The site is down", link)

	} else {
		fmt.Println("The site is up", link)

		//go ReadUrl(response)

	}
}

/*
func ReadUrl(response *http.Response) {
   defer wg.Done()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Could not read contents", err)
	}
	fmt.Printf("The Content length \n%d\n", len(body))
}
*/
