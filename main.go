package main

import (
	"fmt"
	"os"

	chip "github.com/amirulazreen/chip-crawler/src/crawler"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <website_url>")
		os.Exit(1)
	}

	website := os.Args[1]

	result, err := chip.Crawler(website)
	if err != nil {
		fmt.Println(err)

		fmt.Println(result.InputToken, result.OutputToken)
	}
}
