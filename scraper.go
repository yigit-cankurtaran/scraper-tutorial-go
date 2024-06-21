package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

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

		// opening the CSV file
		file, err := os.Create("products.csv")
		if err != nil {
			log.Fatalln("Failed to create output CSV file", err)
		}
		defer file.Close()

		// initializing a file writer
		writer := csv.NewWriter(file)

		// writing the CSV headers
		headers := []string{
			"url",
			"image",
			"name",
			"price",
		}
		writer.Write(headers)

		// writing each product as a CSV row
		for _, product := range products {
			// converting a Product to an array of strings
			record := []string{
				product.url,
				product.image,
				product.name,
				product.price,
			}

			// adding a CSV record to the output file
			writer.Write(record)
		}
		defer writer.Flush()
	})

	c.Visit("https://scrapingcourse.com/ecommerce")
	// Visit() method makes a GET request to the URL passed as an argument
	// and starts the scraping process
	// putting this at the end of the code to start the scraping process

	fmt.Println("Hello, World!")
	fmt.Println(products)

	// timeout error fixed time to export to CSV
}
