package controller

import (
	"fmt"
	"log"
	"strings"

	helper "github.com/amirulazreen/chip-crawler/helper"
	colly "github.com/amirulazreen/chip-crawler/libraries/colly"
	collymodels "github.com/amirulazreen/chip-crawler/libraries/colly/models"
	ai "github.com/amirulazreen/chip-crawler/libraries/together_ai"
	aimodels "github.com/amirulazreen/chip-crawler/libraries/together_ai/models"
	models "github.com/amirulazreen/chip-crawler/src/controller/models"
)

func Controller(website string) (models.WebsiteSummary, error) {
	result := models.WebsiteSummary{}

	scrapedData := colly.CrawlWebsite(website)
	if len(scrapedData) < 1 {
		log.Fatal("error: No link found")
	}

	result.Website = website
	result.Page = scrapedData

	var feeder string
	for i := range scrapedData {
		feeder += scrapedData[i].Content
	}

	prompt := []aimodels.Message{
		instruction,
		{
			Role:    "user",
			Content: fmt.Sprintf("Analyze this webpage:\n\n%s", helper.RemoveDuplicateTexts(feeder)),
		},
	}

	param := aimodels.Request{
		Model:       AIModel,
		Temperature: 0.2,
		Messages:    prompt,
	}

	aiResponse, err := ai.GenerateText(param)
	if err != nil {
		fmt.Printf("Error generating text: %v\n", err)
		return result, err
	}

	urls := make([]string, len(scrapedData))
	for i, page := range scrapedData {
		urls[i] = page.URL
	}

	result.URLS = urls
	result.Content = getContentFromPages(scrapedData)
	result.Summary = aiResponse.Content
	result.InputToken = aiResponse.Usage.PromptTokens
	result.OutputToken = aiResponse.Usage.CompletionTokens

	cost := (result.InputToken * InputCost) + (result.OutputToken * OutputCost)

	// fmt.Println(feeder)
	fmt.Printf("Cost: $ %.6f\n", cost/1_000_000)

	return result, nil
}

func getContentFromPages(pages []collymodels.Page) string {
	var content strings.Builder
	for _, page := range pages {
		content.WriteString(fmt.Sprintf("URL: %s\nTitle: %s\nContent: %s\n\n", page.URL, page.Title, page.Content))
	}
	return content.String()
}
