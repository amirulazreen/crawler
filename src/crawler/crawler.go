package crawler

import (
	"fmt"
	"log"
	"strings"

	helper "github.com/amirulazreen/chip-crawler/helper"
	colly "github.com/amirulazreen/chip-crawler/libraries/colly"
	collymodels "github.com/amirulazreen/chip-crawler/libraries/colly/models"
	ai "github.com/amirulazreen/chip-crawler/libraries/together_ai"
	aimodels "github.com/amirulazreen/chip-crawler/libraries/together_ai/models"
	models "github.com/amirulazreen/chip-crawler/src/crawler/models"
)

func Crawler(website string) (models.WebsiteSummary, error) {
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
		{
			Role: "system",
			// Content: `You are a website content risk analyzer. For each page,
			// 		you will summarize it and rank risk level: high (gambling, alcohol, adult), medium (marketing, crypto), low (safe, educational).
			// 		Return JSON with fields: summary, topics, and risk_level.`,
			Content: `You are a website content risk analyzer. For each page,
					you will summarize it and rank risk level: high (gambling, alcohol, adult), medium (marketing, crypto), low (safe, educational).
					You will tell the known financial, reputation, law, or operation risk.
				 	You will tell whether the company put company name and company number related to Malaysia law (SSM, ROS, SKM),
					You will list out the required documents for due diligence,
					You will make a simple checklist in bullet points for me to check the said documents,
					Return JSON with fields: summary, topics, website type, required_documents, risk_level`,
		},
		{
			Role:    "user",
			Content: fmt.Sprintf("Analyze this webpage:\n\n%s", helper.RemoveDuplicateTexts(feeder)),
		},
	}

	param := aimodels.Request{
		Model:       OpenAIOSS,
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

	cost := (result.InputToken * 0.005) + (result.OutputToken * 0.020)

	fmt.Println(feeder)
	fmt.Printf("Cost: %.2f\n", cost/1_000_000)

	return result, nil
}

// fmt.Println("Saving urls into excel..")

// domain, err := libraries.SaveToExcel(urls, website)
// if err != nil {
// 	log.Fatalf("error: Fail to save %v", err)
// }

// fmt.Printf("✅ All links saved to %s.xlsx\n", domain)

// Helper function to extract content from pages
func getContentFromPages(pages []collymodels.Page) string {
	var content strings.Builder
	for _, page := range pages {
		content.WriteString(fmt.Sprintf("URL: %s\nTitle: %s\nContent: %s\n\n", page.URL, page.Title, page.Content))
	}
	return content.String()
}
