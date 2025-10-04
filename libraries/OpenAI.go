package library

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	models "github.com/amirulazreen/chip-crawler/libraries/models"
)

func CallOpenAI(input string) (*models.OpenAIResponse, error) {
	apiKey := os.Getenv("TOGETHER_AI_KEY")
	url := "https://api.together.xyz/v1/embeddings"

	requestBody := models.OpenAIRequest{
		Model: "openai/gpt-oss-20b",
		Input: input,
		Store: true,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var openAIResp models.OpenAIResponse
	err = json.Unmarshal(body, &openAIResp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return &openAIResp, nil
}

func GenerateText(prompt string) (string, error) {
	apiKey := os.Getenv("TOGETHER_AI_KEY")
	url := "https://api.together.xyz/v1/chat/completions"

	requestBody := models.ChatCompletionRequest{
		Model: "openai/gpt-oss-20b",
		Messages: []models.ChatMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		MaxTokens:   150,
		Temperature: 0.7,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Debug: Print the raw response
	fmt.Printf("Raw API response: %s\n", string(body))

	var chatResp models.ChatCompletionResponse
	err = json.Unmarshal(body, &chatResp)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling response: %v", err)
	}

	// Debug: Print the parsed response structure
	fmt.Printf("Parsed response: %+v\n", chatResp)

	if len(chatResp.Choices) > 0 {
		content := chatResp.Choices[0].Message.Content
		reasoning := chatResp.Choices[0].Message.Reasoning

		fmt.Printf("Generated content: '%s'\n", content)
		fmt.Printf("Generated reasoning: '%s'\n", reasoning)

		// If content is empty but reasoning has text, extract the haiku from reasoning
		if content == "" && reasoning != "" {
			// Look for the haiku pattern in reasoning
			// The haiku appears to be: "Silent circuits hum, / Minds of silicon bloom bright, / Future's quiet code."

			// Try to find the haiku in quotes
			if strings.Contains(reasoning, "Possible haiku:") {
				// Extract text after "Possible haiku:"
				parts := strings.Split(reasoning, "Possible haiku:")
				if len(parts) > 1 {
					haikuText := strings.TrimSpace(parts[1])
					// Remove the quotes and clean up
					haikuText = strings.Trim(haikuText, "\"")
					// Split by / and clean up each line
					lines := strings.Split(haikuText, "/")
					var cleanLines []string
					for _, line := range lines {
						cleanLine := strings.TrimSpace(line)
						cleanLine = strings.Trim(cleanLine, "\"")
						if cleanLine != "" {
							cleanLines = append(cleanLines, cleanLine)
						}
					}
					if len(cleanLines) >= 3 {
						return strings.Join(cleanLines[:3], "\n"), nil
					} else if len(cleanLines) > 0 {
						return strings.Join(cleanLines, "\n"), nil
					}
				}
			}

			// Fallback: look for poetic lines
			lines := strings.Split(reasoning, "\n")
			var haikuLines []string

			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line != "" &&
					!strings.Contains(line, "User:") &&
					!strings.Contains(line, "They want") &&
					!strings.Contains(line, "Provide") &&
					!strings.Contains(line, "Check syllables") &&
					!strings.Contains(line, "Too many") &&
					!strings.Contains(line, "Wait") &&
					!strings.Contains(line, "Need") &&
					!strings.Contains(line, "We need") &&
					!strings.Contains(line, "Haiku is") &&
					!strings.Contains(line, "Let's count") &&
					!strings.Contains(line, "Let's check") &&
					len(line) < 50 { // Haiku lines are typically short
					haikuLines = append(haikuLines, line)
				}
			}

			// Return the first 3 lines that look like a haiku
			if len(haikuLines) >= 3 {
				return strings.Join(haikuLines[:3], "\n"), nil
			} else if len(haikuLines) > 0 {
				return strings.Join(haikuLines, "\n"), nil
			}
		}

		return content, nil
	}

	return "", fmt.Errorf("no response generated")
}
