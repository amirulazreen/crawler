package library

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func InsertURL() string {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <website_url>")
		os.Exit(1)
	}
	return os.Args[1]
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
