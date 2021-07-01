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

	fmt.Println(doc.Find("#zh").Text())

	doc.Find("#zb > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Find("td").Text())
	})

}

func main() {
	ExampleScrape()
}
