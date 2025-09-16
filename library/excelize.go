package library

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/xuri/excelize/v2"
)


func SaveToExcel(pages []Page, website string) {
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.NewSheet(sheet)
	f.SetCellValue(sheet, "A1", "URL")
	f.SetCellValue(sheet, "B1", "Page Title")

	for i, p := range pages {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), p.URL)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), p.Title)
	}

	parsed, err := url.Parse(website)
	if err != nil {
		log.Fatal(err)
	}
	domain := strings.TrimPrefix(parsed.Hostname(), "www.")

	if err := f.SaveAs(domain + ".xlsx"); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("✅ All links saved to %s.xlsx\n", domain)
}
