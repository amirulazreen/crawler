package main

import (
	"fmt"

	"github.com/joho/godotenv"
	// chip "github.com/amirulazreen/chip-crawler/src"
	libraries "github.com/amirulazreen/chip-crawler/libraries"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found, using system environment variables")
	}

	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage: go run main.go <website_url>")
	// 	os.Exit(1)
	// }

	// website := os.Args[1]

	// Option 1: Run the crawler (commented out)
	// chip.Crawler(website)

	// Option 2: Generate text using OpenAI

	text, err := libraries.GenerateText("write a haiku about ai ")
	if err != nil {
		fmt.Printf("Error generating text: %v\n", err)
		return
	}
	fmt.Printf("Generated text: %s\n", text)
}
