package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"nexus/backend/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	db    *gorm.DB
	sqlDB *sql.DB
}

type blogSourceModel struct {
	ID        string
	Name      string
	SiteURL   string
	LogoURL   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (blogSourceModel) TableName() string {
	return "blog_sources"
}

type blogTagModel struct {
	ID   string
	Name string
}

func (blogTagModel) TableName() string {
	return "blog_tags"
}

type blogPostModel struct {
	ID              string
	SourceID        string
	Source          blogSourceModel `gorm:"foreignKey:SourceID"`
	Title           string
	Summary         string
	ExternalURL     string
	HeroImageURL    string
	PublishedAt     time.Time
	ReadTimeMinutes int
	IsFeatured      bool
	Tags            []blogTagModel `gorm:"many2many:blog_post_tags;joinForeignKey:PostID;joinReferences:TagID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (blogPostModel) TableName() string {
	return "blog_posts"
}

type startupTagModel struct {
	ID   string
	Name string
}

func (startupTagModel) TableName() string {
	return "startup_tags"
}

type startupCompanyModel struct {
	ID            string
	Name          string
	Description   string
	Sector        string
	FundingStage  string
	FundingAmount string
	TeamSize      string
	Location      string
	LogoURL       string
	WebsiteURL    string
	CareersURL    string
	IsFeatured    bool
	Tags          []startupTagModel `gorm:"many2many:startup_company_tags;joinForeignKey:CompanyID;joinReferences:TagID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (startupCompanyModel) TableName() string {
	return "startup_companies"
}

func NewStore(ctx context.Context, databaseURL string) (*Store, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, err
	}

	return &Store{
		db:    db,
		sqlDB: sqlDB,
	}, nil
}

func (s *Store) Close() {
	if s.sqlDB != nil {
		_ = s.sqlDB.Close()
	}
}

func (s *Store) ListBlogPosts(ctx context.Context, filters domain.BlogPostFilters) (domain.BlogPostList, error) {
	limit := clamp(filters.Limit, 1, 50, 12)

	query := s.db.WithContext(ctx).
		Model(&blogPostModel{}).
		Preload("Source").
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("blog_tags.id ASC")
		}).
		Scopes(applyBlogFilters(filters)).
		Order("blog_posts.published_at DESC").
		Order("blog_posts.id DESC").
		Limit(limit + 1)

	var records []blogPostModel
	if err := query.Find(&records).Error; err != nil {
		return domain.BlogPostList{}, err
	}

	items := make([]domain.BlogPost, 0, min(len(records), limit))
	for _, record := range records {
		items = append(items, mapBlogPost(record))
	}

	var nextCursor *string
	if len(items) > limit {
		last := items[limit-1]
		cursor := last.PublishedAt.UTC().Format(time.RFC3339) + "::" + last.ID
		nextCursor = &cursor
		items = items[:limit]
	}

	return domain.BlogPostList{
		Items:      items,
		NextCursor: nextCursor,
	}, nil
}

func (s *Store) ListStartups(ctx context.Context, filters domain.StartupFilters) (domain.StartupList, error) {
	limit := clamp(filters.Limit, 1, 50, 20)
	offset := max(filters.Offset, 0)

	countQuery := s.db.WithContext(ctx).
		Model(&startupCompanyModel{}).
		Scopes(applyStartupFilters(filters)).
		Distinct("startup_companies.id")

	var total int64
	if err := countQuery.Count(&total).Error; err != nil {
		return domain.StartupList{}, err
	}

	query := s.db.WithContext(ctx).
		Model(&startupCompanyModel{}).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("startup_tags.id ASC")
		}).
		Scopes(applyStartupFilters(filters)).
		Order("startup_companies.is_featured DESC").
		Order("startup_companies.name ASC").
		Limit(limit).
		Offset(offset)

	var records []startupCompanyModel
	if err := query.Find(&records).Error; err != nil {
		return domain.StartupList{}, err
	}

	items := make([]domain.Startup, 0, len(records))
	for _, record := range records {
		items = append(items, mapStartup(record))
	}

	return domain.StartupList{
		Items:  items,
		Total:  int(total),
		Limit:  limit,
		Offset: offset,
	}, nil
}

func applyBlogFilters(filters domain.BlogPostFilters) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if filters.Query != "" {
			pattern := "%" + filters.Query + "%"
			tx = tx.Where("blog_posts.title ILIKE ? OR blog_posts.summary ILIKE ?", pattern, pattern)
		}

		if filters.SourceID != "" {
			tx = tx.Joins("JOIN blog_sources ON blog_sources.id = blog_posts.source_id").
				Where("blog_sources.id = ?", filters.SourceID)
		}

		if filters.TagID != "" {
			tx = tx.Joins("JOIN blog_post_tags ON blog_post_tags.post_id = blog_posts.id").
				Joins("JOIN blog_tags ON blog_tags.id = blog_post_tags.tag_id").
				Where("blog_tags.id = ?", filters.TagID).
				Select("blog_posts.*").
				Distinct()
		}

		if publishedAt, id, ok := parseCursor(filters.Cursor); ok {
			tx = tx.Where(
				"blog_posts.published_at < ? OR (blog_posts.published_at = ? AND blog_posts.id < ?)",
				publishedAt,
				publishedAt,
				id,
			)
		}

		return tx
	}
}

func applyStartupFilters(filters domain.StartupFilters) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if filters.Query != "" {
			pattern := "%" + filters.Query + "%"
			tx = tx.Where("startup_companies.name ILIKE ? OR startup_companies.description ILIKE ?", pattern, pattern)
		}

		if filters.Sector != "" {
			tx = tx.Where("startup_companies.sector ILIKE ?", filters.Sector)
		}

		if filters.Location != "" {
			tx = tx.Where("startup_companies.location ILIKE ?", "%"+filters.Location+"%")
		}

		if filters.Stage != "" {
			tx = tx.Where("startup_companies.funding_stage ILIKE ?", filters.Stage)
		}

		if filters.TagID != "" {
			tx = tx.Joins("JOIN startup_company_tags ON startup_company_tags.company_id = startup_companies.id").
				Joins("JOIN startup_tags ON startup_tags.id = startup_company_tags.tag_id").
				Where("startup_tags.id = ?", filters.TagID).
				Select("startup_companies.*").
				Distinct()
		}

		return tx
	}
}

func mapBlogPost(record blogPostModel) domain.BlogPost {
	return domain.BlogPost{
		ID:              record.ID,
		Title:           record.Title,
		Summary:         record.Summary,
		ExternalURL:     record.ExternalURL,
		HeroImageURL:    record.HeroImageURL,
		PublishedAt:     record.PublishedAt,
		ReadTimeMinutes: record.ReadTimeMinutes,
		Source: domain.BlogSource{
			ID:      record.Source.ID,
			Name:    record.Source.Name,
			SiteURL: record.Source.SiteURL,
		},
		Tags: mapBlogTags(record.Tags),
	}
}

func mapStartup(record startupCompanyModel) domain.Startup {
	return domain.Startup{
		ID:            record.ID,
		Name:          record.Name,
		Description:   record.Description,
		Sector:        record.Sector,
		FundingStage:  record.FundingStage,
		FundingAmount: record.FundingAmount,
		TeamSize:      record.TeamSize,
		Location:      record.Location,
		LogoURL:       record.LogoURL,
		WebsiteURL:    record.WebsiteURL,
		CareersURL:    record.CareersURL,
		Tags:          mapStartupTags(record.Tags),
	}
}

func mapBlogTags(tags []blogTagModel) []string {
	values := make([]string, 0, len(tags))
	for _, tag := range tags {
		values = append(values, tag.ID)
	}
	return values
}

func mapStartupTags(tags []startupTagModel) []string {
	values := make([]string, 0, len(tags))
	for _, tag := range tags {
		values = append(values, tag.ID)
	}
	return values
}

func parseCursor(cursor string) (time.Time, string, bool) {
	if cursor == "" {
		return time.Time{}, "", false
	}

	parts := strings.Split(cursor, "::")
	if len(parts) != 2 {
		return time.Time{}, "", false
	}

	publishedAt, err := time.Parse(time.RFC3339, parts[0])
	if err != nil {
		return time.Time{}, "", false
	}

	return publishedAt, parts[1], true
}

func clamp(value int, minValue int, maxValue int, fallback int) int {
	if value == 0 {
		return fallback
	}
	if value < minValue {
		return minValue
	}
	if value > maxValue {
		return maxValue
	}
	return value
}

func max(value int, minimum int) int {
	if value < minimum {
		return minimum
	}
	return value
}

func min(value int, maximum int) int {
	if value < maximum {
		return value
	}
	return maximum
}
