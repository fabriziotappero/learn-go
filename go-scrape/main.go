package main

import (
	"fmt"
	"log"
	"github.com/gocolly/colly/v2"
)

func main() {
	// Create a new collector
	c := colly.NewCollector()

	// Set up a callback for each quote element
	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		quoteText := e.ChildText("span.text")
		quoteAuthor := e.ChildText("span small.author")
		quoteTags := e.ChildAttr("div.tags a.tag", "href")

		fmt.Printf("Quote: %s\nAuthor: %s\nTags: %s\n\n", quoteText, quoteAuthor, quoteTags)
	})

	// Set up an error callback
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Request URL: %s failed with response: %v\n", r.Request.URL, r)
	})

	// Start the scraping
	err := c.Visit("http://quotes.toscrape.com")
	if err != nil {
		log.Fatal(err)
	}
}
