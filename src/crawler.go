package src

import (
	"fmt"
	"log"

	libraries "github.com/amirulazreen/chip-crawler/libraries"
)

func Crawler(website string) {
	fmt.Println("Crawling: ", website)

	urls := libraries.CrawlWebsite(website)
	if len(urls) < 1 {
		log.Fatal("error: No link found")
	}

	fmt.Println("Saving urls into excel..")

	domain, err := libraries.SaveToExcel(urls, website)
	if err != nil {
		log.Fatalf("error: Fail to save %v", err)
	}

	fmt.Printf("✅ All links saved to %s.xlsx\n", domain)
}
