package colly

import (
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	models "github.com/amirulazreen/chip-crawler/libraries/colly/models"

	"github.com/gocolly/colly/v2"
)

func CrawlWebsite(website string) []models.Page {
	var results []models.Page
	seen := make(map[string]bool)

	fmt.Println("Found links")
	u, err := url.Parse(website)
	if err != nil {
		log.Fatal(err)
	}
	host := u.Hostname()

	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:89.0) Gecko/20100101 Firefox/89.0",
	}

	c := colly.NewCollector(
		colly.AllowedDomains(host),
		colly.MaxDepth(2),
		colly.UserAgent(userAgents[rand.Intn(len(userAgents))]),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*." + host,
		Parallelism: 10,
		Delay:       200 * time.Millisecond,
		RandomDelay: 300 * time.Millisecond,
	})

	c.WithTransport(&http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
		IdleConnTimeout:       30 * time.Second,
		DisableKeepAlives:     false,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})

	c.SetRequestTimeout(10 * time.Second)

	c.OnHTML("title", func(e *colly.HTMLElement) {
		url := e.Request.URL.String()
		if !seen[url] {
			seen[url] = true
			results = append(results, models.Page{
				URL:     url,
				Title:   e.Text,
				Content: "",
			})
		} else {
			for i := range results {
				if results[i].URL == url && results[i].Title == "" {
					results[i].Title = e.Text
				}
			}
		}
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		url := e.Request.URL.String()

		content := strings.TrimSpace(e.Text)
		content = strings.ReplaceAll(content, "\n\n\n", "\n\n")
		content = strings.ReplaceAll(content, "  ", " ")
		content = strings.TrimSpace(content)

		for i := range results {
			if results[i].URL == url {
				results[i].Content = content
				break
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.5")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("DNT", "1")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link == "" {
			return
		}

		if strings.HasPrefix(link, "mailto:") ||
			strings.HasPrefix(link, "tel:") ||
			strings.HasPrefix(link, "javascript:") ||
			strings.Contains(link, "#") ||
			strings.HasSuffix(link, ".pdf") ||
			strings.HasSuffix(link, ".jpg") ||
			strings.HasSuffix(link, ".png") ||
			strings.HasSuffix(link, ".gif") {
			return
		}

		if !seen[link] {
			seen[link] = true
			results = append(results, models.Page{
				URL:     link,
				Title:   "",
				Content: "",
			})
			fmt.Println(link)
		}

		if linkURL, err := url.Parse(link); err == nil {
			if strings.HasSuffix(linkURL.Hostname(), host) {
				e.Request.Visit(link)
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		if len(results) > 1000 {
			r.Abort()
		}
	})

	c.Visit(website)
	return results
}
