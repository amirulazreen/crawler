package library

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func CrawlWebsite(website string) []Page {
	var results []Page

	c := colly.NewCollector(
		colly.MaxDepth(4),
	)

	c.OnHTML("title", func(e *colly.HTMLElement) {
		results = append(results, Page{
			URL:   e.Request.URL.String(),
			Title: e.Text,
		})
		fmt.Printf("Saved: %s | %s\n", e.Request.URL.String(), e.Text)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" {
			fmt.Println("Found link:", link)
			e.Request.Visit(link)
		}
	})

	c.Visit(website)
	return results
}
