package main

import (
	"os"

	chip "github.com/amirulazreen/chip-crawler/src"
)

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage: go run main.go <website_url>")
	// 	return
	// }

	// website := os.Args[1]
	// fmt.Println("🔗 Crawling:", website)

	// f := excelize.NewFile()
	// sheet := "Sheet1"
	// f.NewSheet(sheet)
	// f.SetCellValue(sheet, "A1", "URL")
	// f.SetCellValue(sheet, "B1", "Page Title")

	// row := 2

	// c := colly.NewCollector(
	// 	colly.MaxDepth(4),
	// )


	// c.OnHTML("title", func(e *colly.HTMLElement) {
	// 	link := e.Request.URL.String()
	// 	title := e.Text

	// 	f.SetCellValue(sheet, fmt.Sprintf("A%d", row), link)
	// 	f.SetCellValue(sheet, fmt.Sprintf("B%d", row), title)
	// 	fmt.Printf("Saved: %s | %s\n", link, title)
	// 	row++
	// })

	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	link := e.Request.AbsoluteURL(e.Attr("href"))
	// 	if link != "" {
	// 		fmt.Println("Found link:", link)
	// 		e.Request.Visit(link)
	// 	}
	// })

	// c.Visit(website)

	// parsed, err := url.Parse(website)
	// if err != nil {
	// 	panic(err)
	// }
	// host := parsed.Hostname()
	// domain := strings.TrimPrefix(host, "www.")

	// if err := f.SaveAs(domain + ".xlsx"); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("✅ All links saved to %s.xlsx\n", domain)
	//

	if len(os.Args) < 2 {
		os.Exit(1)
	}

	website := os.Args[1]

	chip.Crawler(website)
}
