package library

import (
	"fmt"

	helper "github.com/amirulazreen/chip-crawler/helper"
	models "github.com/amirulazreen/chip-crawler/libraries/models"

	"github.com/xuri/excelize/v2"
)


func SaveToExcel(pages []models.Page, website string) (string,error){
	f := excelize.NewFile()
	sheet := "sheet1"
	f.NewSheet(sheet)
	f.SetCellValue(sheet, "A1", "URL")
	f.SetCellValue(sheet, "B1", "Page Title")

	for i, p := range pages {
		row := i + 2
		url := helper.Sanitize(p.URL)
		title := helper.Sanitize(p.Title)
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), url)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), title)
	}

	f.SetColWidth(sheet, "A", "A", 100)
	f.SetColWidth(sheet, "B", "B", 100)

	domain, err := helper.GetDomainName(website)
	if err != nil{
		return "", err
	}

	if err := f.SaveAs(domain + ".xlsx"); err != nil {
		return "", err
	}
	return domain, nil
}
