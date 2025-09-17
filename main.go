package main

import (
	"os"

	chip "github.com/amirulazreen/chip-crawler/src"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	website := os.Args[1]

	chip.Crawler(website)
}
