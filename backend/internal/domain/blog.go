package domain

import "time"

type BlogSource struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	SiteURL string `json:"siteUrl"`
}

type BlogPost struct {
	ID              string     `json:"id"`
	Title           string     `json:"title"`
	Summary         string     `json:"summary"`
	ExternalURL     string     `json:"externalUrl"`
	HeroImageURL    string     `json:"heroImageUrl"`
	PublishedAt     time.Time  `json:"publishedAt"`
	ReadTimeMinutes int        `json:"readTimeMinutes"`
	Source          BlogSource `json:"source"`
	Tags            []string   `json:"tags"`
}

type BlogPostList struct {
	Items      []BlogPost `json:"items"`
	NextCursor *string    `json:"nextCursor"`
}

type BlogPostFilters struct {
	Cursor   string
	Limit    int
	Query    string
	SourceID string
	TagID    string
}
