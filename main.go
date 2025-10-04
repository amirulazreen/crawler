package main

import (
	"fmt"

	// chip "github.com/amirulazreen/chip-crawler/src"
	libraries "github.com/amirulazreen/chip-crawler/libraries"
	models "github.com/amirulazreen/chip-crawler/libraries/models"
)

func main() {

	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage: go run main.go <website_url>")
	// 	os.Exit(1)
	// }

	// website := os.Args[1]

	// Option 1: Run the crawler (commented out)
	// chip.Crawler(website)

	// Option 2: Generate text using OpenAI

	param := models.Request{
		Model: "openai/gpt-oss-20b",
		Messages: []models.Message{
			{Role: "user", Content: "write a haiku about ai with 3 lines"},
		},
	}

	text, err := libraries.GenerateText(param)
	if err != nil {
		fmt.Printf("Error generating text: %v\n", err)
		return
	}
	fmt.Printf("Generated text: %s\n", text)
}
