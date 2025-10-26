package togetherai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/amirulazreen/chip-crawler/libraries/together_ai/models"
)

func GenerateText(param models.Request) (models.SummarizeData, error) {
	result := models.SummarizeData{}

	if param.APIKey == "" {
		return result, fmt.Errorf("missing Together AI API key")
	}

	reqBody := models.Request{
		Model:       param.Model,
		Temperature: param.Temperature,
		Messages:    param.Messages,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return result, fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.together.xyz/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return result, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+param.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("API returned status: %s", resp.Status)
	}

	var res models.Response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return result, fmt.Errorf("failed to decode response: %v", err)
	}

	if len(res.Choices) == 0 {
		return result, fmt.Errorf("no choices returned from API")
	}

	result.Content = res.Choices[0].Message.Content
	result.Usage.PromptTokens = res.Usage.PromptTokens
	result.Usage.CompletionTokens = res.Usage.CompletionTokens

	return result, nil
}
