package library

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	models "github.com/amirulazreen/chip-crawler/libraries/models"
	"github.com/joho/godotenv"
)

func GenerateText(param models.Request) (string, error) {
	inputToken := 0
	outputToken := 0

	godotenv.Load()
	apiKey := os.Getenv("TOGETHER_AI_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("missing API key")
	}

	reqBody := models.Request{
		Model:       param.Model,
		Temperature: param.Temperature,
		Messages:    param.Messages,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.together.xyz/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned status: %s", resp.Status)
	}

	var res models.Response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	if len(res.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from API")
	}

	inputToken = res.Usage.PromptTokens
	outputToken = res.Usage.CompletionTokens

	fmt.Println(inputToken)
	fmt.Println(outputToken)

	return res.Choices[0].Message.Content, nil
}
