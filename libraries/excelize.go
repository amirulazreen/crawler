package library

import (
	"fmt"

	helper "github.com/amirulazreen/chip-crawler/helper"
	models "github.com/amirulazreen/chip-crawler/libraries/models"

	"github.com/xuri/excelize/v2"
)

func SaveToExcel(pages []models.Page, website string) (string, error) {
	f := excelize.NewFile()
	sheet := "sheet1"
	f.NewSheet(sheet)
	f.SetCellValue(sheet, "A1", "URL")
	f.SetCellValue(sheet, "B1", "Page Title")
	f.SetCellValue(sheet, "C1", "Page Content")
	for i, p := range pages {
		row := i + 2
		url := helper.Sanitize(p.URL)
		title := helper.Sanitize(p.Title)
		content := helper.Sanitize(p.Content)
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), url)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), title)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), content)
	}

	f.SetColWidth(sheet, "A", "A", 80)
	f.SetColWidth(sheet, "B", "B", 80)
	f.SetColWidth(sheet, "C", "C", 80)

	domain, err := helper.GetDomainName(website)
	if err != nil {
		return "", err
	}

	if err := f.SaveAs(domain + ".xlsx"); err != nil {
		return "", err
	}
	return domain, nil
}
