package controller

import (
	"fmt"

	helper "github.com/amirulazreen/chip-crawler/helper"
	"github.com/amirulazreen/chip-crawler/libraries/whois"
	whoismodels "github.com/amirulazreen/chip-crawler/libraries/whois/models"
	models "github.com/amirulazreen/chip-crawler/src/controller/models"
)

func Controller(website string) (models.WebsiteSummary, error) {
	result := models.WebsiteSummary{}

	domain, err := helper.GetDomainName(website)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	param := whoismodels.WhoIsRequest{
		Website: domain,
	}

	test, err := whois.GetWhoisData(param)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Println(test)

	return result, nil
}
