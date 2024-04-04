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

/*
// what to do for dynamic pages

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
*/
