package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://rates.goldenchennai.com/gold-rate/chennai-gold-rate-today/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc.Find("#wp > main > h1").First().Text())
	/*
		// Find the review items
		doc.Find(".question-summary .summary").Each(func(i int, s *goquery.Selection) {
			title := s.Find("H3").Text()
			fmt.Println(i, title)
		})
	*/

}

func main() {
	ExampleScrape()
}
