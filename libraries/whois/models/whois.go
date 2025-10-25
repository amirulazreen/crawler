package models

type WhoIsRequest struct {
	Website string
	APIKey  string
}

type WhoisResponse struct {
	WhoisRecord WhoisRecord `json:"WhoisRecord"`
}

type WhoisRecord struct {
	CreatedDate           string       `json:"createdDate"`
	UpdatedDate           string       `json:"updatedDate"`
	ExpiresDate           string       `json:"expiresDate"`
	Registrant            Contact      `json:"registrant"`
	TechnicalContact      Contact      `json:"technicalContact"`
	DomainName            string       `json:"domainName"`
	NameServers           NameServers  `json:"nameServers"`
	Status                string       `json:"status"`
	RawText               string       `json:"rawText"`
	ParseCode             int          `json:"parseCode"`
	Header                string       `json:"header"`
	StrippedText          string       `json:"strippedText"`
	Footer                string       `json:"footer"`
	Audit                 Audit        `json:"audit"`
	RegistrarName         string       `json:"registrarName"`
	RegistrarIANAID       string       `json:"registrarIANAID"`
	CreatedDateNormalized string       `json:"createdDateNormalized"`
	UpdatedDateNormalized string       `json:"updatedDateNormalized"`
	ExpiresDateNormalized string       `json:"expiresDateNormalized"`
	RegistryData          RegistryData `json:"registryData"`
	DomainAvailability    string       `json:"domainAvailability"`
	ContactEmail          string       `json:"contactEmail"`
	DomainNameExt         string       `json:"domainNameExt"`
	EstimatedDomainAge    int          `json:"estimatedDomainAge"`
}

type Contact struct {
	Organization string `json:"organization,omitempty"`
	Country      string `json:"country"`
	CountryCode  string `json:"countryCode"`
	RawText      string `json:"rawText"`
}

type NameServers struct {
	RawText   string   `json:"rawText"`
	HostNames []string `json:"hostNames"`
	IPs       []string `json:"ips"`
}

type Audit struct {
	CreatedDate string `json:"createdDate"`
	UpdatedDate string `json:"updatedDate"`
}

type RegistryData struct {
	CreatedDate           string      `json:"createdDate"`
	UpdatedDate           string      `json:"updatedDate"`
	ExpiresDate           string      `json:"expiresDate"`
	DomainName            string      `json:"domainName"`
	NameServers           NameServers `json:"nameServers"`
	Status                string      `json:"status"`
	RawText               string      `json:"rawText"`
	ParseCode             int         `json:"parseCode"`
	Header                string      `json:"header"`
	StrippedText          string      `json:"strippedText"`
	Footer                string      `json:"footer"`
	Audit                 Audit       `json:"audit"`
	RegistrarName         string      `json:"registrarName"`
	RegistrarIANAID       string      `json:"registrarIANAID"`
	CreatedDateNormalized string      `json:"createdDateNormalized"`
	UpdatedDateNormalized string      `json:"updatedDateNormalized"`
	ExpiresDateNormalized string      `json:"expiresDateNormalized"`
	WhoisServer           string      `json:"whoisServer"`
}
