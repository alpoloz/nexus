package domain

type Startup struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Sector        string   `json:"sector"`
	FundingStage  string   `json:"fundingStage"`
	FundingAmount string   `json:"fundingAmount"`
	TeamSize      string   `json:"teamSize"`
	Location      string   `json:"location"`
	LogoURL       string   `json:"logoUrl"`
	WebsiteURL    string   `json:"websiteUrl"`
	CareersURL    string   `json:"careersUrl"`
	Tags          []string `json:"tags"`
}

type StartupList struct {
	Items  []Startup `json:"items"`
	Total  int       `json:"total"`
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
}

type StartupFilters struct {
	Limit    int
	Offset   int
	Query    string
	Sector   string
	Location string
	Stage    string
	TagID    string
}
