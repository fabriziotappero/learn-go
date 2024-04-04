package main

import (
	"fmt"
	"log"
	"github.com/gocolly/colly/v2"
)

func fetchHTML(url string) (string, error) {
    // Initialize a new browser context
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    // Navigate to the URL and fetch the rendered HTML
    var htmlContent string
    err := chromedp.Run(ctx,
        chromedp.Navigate(url),
        chromedp.OuterHTML("html", &htmlContent),
    )
    if err != nil {
        return "", err
    }
  
    return htmlContent, nil
}

func main() {
    // Fetch the rendered HTML content
    htmlContent, err := fetchHTML("http://quotes.toscrape.com/js")
    if err != nil {
        log.Fatal(err)
    }

    // Create a new collector
    c := colly.NewCollector()

    // Set up the callbacks as before
    // ...

    // Use the parsed HTML content as the input for the scraper
    err = c.ParseString(htmlContent)
    if err != nil {
        log.Fatal(err)
    }
}
