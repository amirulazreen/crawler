package whois

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	models "github.com/amirulazreen/chip-crawler/libraries/whois/models"
)

func GetWhoisData(param models.WhoIsRequest) (models.WhoisPartialResult, error) {
	result := models.WhoisPartialResult{}

	if param.APIKey == "" {
		return result, fmt.Errorf("missing WHOIS API key")
	}

	url := fmt.Sprintf(
		"https://www.whoisxmlapi.com/whoisserver/WhoisService?outputFormat=JSON&domainName=%s&apiKey=%s",
		param.Website,
		param.APIKey,
	)
	req, _ := http.NewRequest("GET", url, nil)

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

	var partial models.WhoisPartialResponse
	if err := json.Unmarshal(body, &partial); err != nil {
		return result, fmt.Errorf("failed to parse response: %w", err)
	}

	AvgMonth := 30.44
	estimatedMonth := float64(partial.WhoisRecord.EstimatedDomainAge) / AvgMonth

	result.DomainName = partial.WhoisRecord.DomainName
	result.CreatedDate = "Created Date: " + partial.WhoisRecord.CreatedDate.Format("2006-01-02")
	result.Country = "Country: " + partial.WhoisRecord.Registrant.Country
	result.EstimatedDomainAge = "Estimated Age: " + strconv.Itoa(int(estimatedMonth)) + " months"

	return result, nil
}
