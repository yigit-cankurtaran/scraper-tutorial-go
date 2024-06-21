package main

import (
	"fmt"
	// import colly
	"github.com/gocolly/colly"
)

func main() {
	// scraping logic here
	// gonna use ScrapingCourse.com as a test

	c := colly.NewCollector()
	// colly's main entity is Collector
	// allows for HTTP requests and data extraction
	c.OnRequest(func(r *colly.Request) {
		// attaching callback functions to the Collector
		fmt.Println("Visiting", r.URL)
	})
	// OnRequest() method is called before each request
	// it prints the URL of the page being visited
	c.OnError(func(r *colly.Response, err error) {
		// OnError() method is called when an error occurs
		// this is response request is request
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnResponse(func(r *colly.Response) {
		// OnResponse() method is called after a response is received
		fmt.Println("Visited", r.Request.URL)
	})
	c.OnHTML("a", func(e *colly.HTMLElement) {
		// OnHTML() method is called when an HTML element is found
		// it prints the URL of the page being visited
		link := e.Attr("href")
		fmt.Println("Link found:", link)
	})
	c.OnScraped(func(r *colly.Response) {
		// OnScraped() method is called after the scraping is done
		fmt.Println("Finished scraping ", r.Request.URL)
	})

	c.Visit("https://scrapingcourse.com/")
	// Visit() method makes a GET request to the URL passed as an argument
	// and starts the scraping process
	// putting this at the end of the code to start the scraping process

	fmt.Println("Hello, World!")
}
