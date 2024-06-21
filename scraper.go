package main

import (
	"fmt"
	// import colly
	"github.com/gocolly/colly"
)

type Product struct {
	url, image, name, price string
}

func main() {
	// scraping logic here
	// gonna use ScrapingCourse.com as a test

	var products []Product

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
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		// OnHTML() method is called when an HTML element is found
		// the first argument is the HTML element to look for
		// we can see that on the console before we start scraping

		product := Product{}

		// the data of our interest
		product.url = e.ChildAttr("a", "href")
		product.image = e.ChildAttr("img", "src")
		product.name = e.ChildText("h2")
		product.price = e.ChildText(".price")

		products = append(products, product)

	})
	c.OnScraped(func(r *colly.Response) {
		// OnScraped() method is called after the scraping is done
		fmt.Println("Finished scraping ", r.Request.URL)
	})

	c.Visit("https://scrapingcourse.com/ecommerce")
	// Visit() method makes a GET request to the URL passed as an argument
	// and starts the scraping process
	// putting this at the end of the code to start the scraping process

	fmt.Println("Hello, World!")
	fmt.Println(products)
}
