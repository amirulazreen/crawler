package library

import (
	"net/url"
	"strings"
)

func Sanitize(s string) string {
	s = strings.TrimSpace(s)

	if strings.HasPrefix(s, "=") || strings.HasPrefix(s, "+") ||
		strings.HasPrefix(s, "-") || strings.HasPrefix(s, "@") {
		s = "'" + s
	}
	return s
}

func GetDomainName(website string) (string, error) {
	parsed, err := url.Parse(website)
	if err != nil {
		return "", err
	}

	host := parsed.Hostname()
	domain := strings.TrimPrefix(host, "www.")

	return domain, nil
}

func RemoveDuplicateTexts(input string) string {
	// Split the string into words (you can adjust to split by line or punctuation)
	words := strings.Fields(input)

	seen := make(map[string]bool)
	var result []string

	for _, word := range words {
		normalized := strings.ToLower(strings.TrimSpace(word))
		if normalized == "" {
			continue
		}

		if !seen[normalized] {
			seen[normalized] = true
			result = append(result, word) // keep original case
		}
	}

	return strings.Join(result, " ")
}
