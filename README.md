Some curated material to learn Go.
<img align="right" width="200" src="gopher.png">


### Scrape static and dynamic sites
https://reintech.io/blog/building-web-scrapers-go-colly

    go mod init go-scrape
    go get -u github.com/gocolly/colly/v2
    go get -u github.com/chromedp/chromedp
    go mod tidy
    go run 
    go build


