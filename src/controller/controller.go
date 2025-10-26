package controller

import (
	"fmt"

	helper "github.com/amirulazreen/chip-crawler/helper"
	colly "github.com/amirulazreen/chip-crawler/libraries/colly"
	ai "github.com/amirulazreen/chip-crawler/libraries/together_ai"
	aimodels "github.com/amirulazreen/chip-crawler/libraries/together_ai/models"
	whois "github.com/amirulazreen/chip-crawler/libraries/whois"
	whoismodels "github.com/amirulazreen/chip-crawler/libraries/whois/models"
	models "github.com/amirulazreen/chip-crawler/src/controller/models"
)

func Controller(param models.Configs) (models.WebsiteSummary, error) {
	result := models.WebsiteSummary{}

	collyResponse := colly.CrawlWebsite(param.Website)
	if len(collyResponse) < 1 {
		return result, fmt.Errorf("error:colly:no link found")
	}

	var scrapedHTML string
	for i := range collyResponse {
		scrapedHTML += collyResponse[i].Content
	}

	prompt := []aimodels.Message{
		instruction,
		{
			Role:    "user",
			Content: fmt.Sprintf("Analyze this webpage:\n\n%s", helper.RemoveDuplicateTexts(scrapedHTML)),
		},
	}

	aiParam := aimodels.Request{
		APIKey:      param.TogetherAIPIKey,
		Model:       AIModel,
		Temperature: 0.2,
		Messages:    prompt,
	}

	aiResponse, err := ai.GenerateText(aiParam)
	if err != nil {
		return result, fmt.Errorf("error:togetherAI:%v", err)
	}

	urls := make([]string, len(collyResponse))
	for i, page := range collyResponse {
		urls[i] = page.URL
	}

	whoisParam := whoismodels.WhoIsRequest{
		APIKey:  param.WhoisAPIKey,
		Website: param.Website,
	}

	whoisResponse, err := whois.GetWhoisData(whoisParam)
	if err != nil {
		return result, fmt.Errorf("error:whois:%v", err)
	}

	result.Website = param.Website
	result.Page = collyResponse
	result.URLS = urls
	result.Content = helper.GetContentFromPages(collyResponse)
	result.Summary = aiResponse.Content
	result.InputToken = aiResponse.Usage.PromptTokens
	result.OutputToken = aiResponse.Usage.CompletionTokens
	result.TotalCost = (result.InputToken * InputCost) + (result.OutputToken * OutputCost)
	result.WhoisResult = whoisResponse

	return result, nil
}
