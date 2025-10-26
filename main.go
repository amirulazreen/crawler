package main

import (
	"fmt"
	"os"

	chip "github.com/amirulazreen/chip-crawler/src/controller"
	"github.com/amirulazreen/chip-crawler/src/controller/models"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <website_url>")
		os.Exit(1)
	}

	godotenv.Load()
	param := models.Configs{
		Website:         os.Args[1],
		TogetherAIPIKey: os.Getenv("TOGETHER_AI_KEY"),
		WhoisAPIKey:     os.Getenv("WHOIS_KEY"),
	}

	result, err := chip.Controller(param)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result.Summary)
	fmt.Println("\nResult from Whois")
	fmt.Println(result.WhoisResult.DomainName)
	fmt.Println(result.WhoisResult.Country)
	fmt.Println(result.WhoisResult.CreatedDate)
	fmt.Println(result.WhoisResult.EstimatedDomainAge)

	fmt.Printf("\nCost per crawl: $ %.6f\n", result.TotalCost/1_000_000)
}
