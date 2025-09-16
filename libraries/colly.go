package library

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	models "github.com/amirulazreen/chip-crawler/libraries/models"

	"github.com/gocolly/colly/v2"
)

func CrawlWebsite(website string) []models.Page {
	var results []models.Page
	seen := make(map[string]bool) // track saved links

	u, err := url.Parse(website)
	if err != nil {
		log.Fatal(err)
	}
	host := u.Hostname()

	c := colly.NewCollector(
		colly.AllowedDomains(host),
		colly.MaxDepth(0), // unlimited depth (only for domain/subdomains)
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*." + host,
		Parallelism: 5,
		Delay:       1 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	c.WithTransport(&http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
		IdleConnTimeout:   10 * time.Second,
		DisableKeepAlives: false,
	})

	c.SetRequestTimeout(15 * time.Second)

	// When saving page title
	c.OnHTML("title", func(e *colly.HTMLElement) {
	    url := e.Request.URL.String()
	    if !seen[url] {
	        seen[url] = true
	        results = append(results, models.Page{
	            URL:   url,
	            Title: e.Text, // actual page title
	        })
	        fmt.Printf("Saved: %s | %s\n", url, e.Text)
	    } else {
	        // If already saved before (e.g. as a link), update title
	        for i := range results {
	            if results[i].URL == url && results[i].Title == "" {
	                results[i].Title = e.Text
	                fmt.Printf("Updated title for: %s | %s\n", url, e.Text)
	            }
	        }
	    }
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
    link := e.Request.AbsoluteURL(e.Attr("href"))
    if link == "" {
        return
    }

    if !seen[link] {
        seen[link] = true
        results = append(results, models.Page{
            URL:   link,
            Title: "",
        })
        fmt.Println("Found link:", link)
    }

    if linkURL, err := url.Parse(link); err == nil {
	        if strings.HasSuffix(linkURL.Hostname(), host) {
	            e.Request.Visit(link)
	        }
	    }
	})


	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", r.Request.URL, err)
	})

	c.Visit(website)
	return results
}
