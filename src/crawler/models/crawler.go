package models

import (
	colly "github.com/amirulazreen/chip-crawler/libraries/colly/models"
)

type Page = colly.Page

type WebsiteSummary struct {
	Website     string
	URLS        []string
	Content     string
	Summary     string
	Topic       []string
	RiskLevel   string
	InputToken  int
	OutputToken int
	Page        []colly.Page `json:"page"`
}
