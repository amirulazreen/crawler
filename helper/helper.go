package library

import (
	"net/url"
	"os"
	"strings"
)

func InsertURL() string {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	return os.Args[1]
}

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
