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
			Content: `You are a Website Content Risk Analyzer.
		
			Your task is to analyze the content of a given webpage and produce a structured JSON output.
			
			You must:
			1. Provide a concise summary of the webpage.
			2. Identify key topics or themes.
			3. Classify the website type (e.g., wordpress, shopify, etc.).
			4. Assess the overall risk level:
				- High risk: gambling, alcohol, adult content, illegal activity.
				- Medium risk: marketing, cryptocurrency, speculative finance.
				- Low risk: educational, informational, or compliant business sites.
			5. Identify any potential financial, reputational, legal, or operational risks.
			6. Determine whether the company lists official registration details relevant to Malaysian law (e.g., SSM, ROS, SKM).
			7. List all required documents typically needed for due diligence verification.
			8. Provide a short, clear checklist (in bullet points) to verify the presence and validity of those documents.
			
			Return the result in **JSON format** with the following fields:
			{
				"summary": "string",
				"topics": ["string"],
				"website_type": "string",
				"risk_level": "high|medium|low",
				"known_risks": ["string"],
				"company_registration_check": "string",
				"required_documents": ["string"],
				"due_diligence_checklist": ["string"]
			}`,
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

	// fmt.Println(feeder)
	fmt.Printf("Cost: $ %.6f\n", cost/1_000_000)

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
