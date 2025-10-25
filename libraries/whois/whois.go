package whois

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	models "github.com/amirulazreen/chip-crawler/libraries/whois/models"
	"github.com/joho/godotenv"
)

func GetWhoisData(param models.WhoIsRequest) (models.WhoisResponse, error) {
	result := models.WhoisResponse{}

	godotenv.Load()
	apiKey := os.Getenv("WHOIS_KEY")
	if apiKey == "" {
		return result, fmt.Errorf("missing WHOIS API key")
	}

	param.APIKey = apiKey

	jsonData, err := json.Marshal(param)
	if err != nil {
		return result, fmt.Errorf("failed to marshal json: %w", err)
	}

	req, err := http.NewRequest("POST", "https://www.whoisxmlapi.com/whoisserver/WhoisService", bytes.NewBuffer(jsonData))
	if err != nil {
		return result, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("failed to read response: %w", err)
	}

	fmt.Println(&body)

	return result, nil
}
