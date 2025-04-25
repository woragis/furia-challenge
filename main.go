package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector()
	url := "https://x.com/furia"


	c.OnResponse(func(r *colly.Response) {
		doc, err := htmlquery.Parse(strings.NewReader(string(r.Body)))
		if err != nil {
			log.Fatal(err)
		}

		// XPath to find the followers link
		node := htmlquery.FindOne(doc, "//a[contains(@href, '/verified_followers')]")
		if node == nil {
			fmt.Println("No followers link found.")
			return
		}

		href := htmlquery.SelectAttr(node, "href")
		text := htmlquery.InnerText(node)

		fmt.Printf("Text: %s\nLink: %s\n", text, href)
	})
	


	c.Visit(url)
}
