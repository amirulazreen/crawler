package models

import (
	colly "github.com/amirulazreen/chip-crawler/libraries/colly/models"
	"github.com/amirulazreen/chip-crawler/libraries/whois/models"
)

type Page = colly.Page

type Configs struct {
	Website         string
	TogetherAIPIKey string
	WhoisAPIKey     string
}

type WebsiteSummary struct {
	Website     string
	URLS        []string
	Content     string
	Summary     string
	Topic       []string
	RiskLevel   string
	InputToken  float64
	OutputToken float64
	TotalCost   float64
	Page        []colly.Page `json:"page"`
	WhoisResult models.WhoisPartialResult
}
