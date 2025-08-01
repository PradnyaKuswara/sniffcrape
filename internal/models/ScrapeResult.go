package models

import (
	"time"

	"gorm.io/gorm"
)

type ScrapeResult struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Url         string         `json:"url"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type ScrapeResultRequest struct {
	Url         string `json:"url" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type ScrapeResultRequestOnlyUrl struct {
	Url string `json:"url" validate:"required"`
}

type ScrapeResultResponse struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	Canonical   string `json:"canonical"`
	Robots      string `json:"robots"`
	Viewport    string `json:"viewport"`
	Charset     string `json:"charset"`
	Author      string `json:"author"`

	OgTitle string `json:"og_title"`
	OgDesc  string `json:"og_description"`
	OgImage string `json:"og_image"`
	OgUrl   string `json:"og_url"`

	TwTitle string `json:"twitter_title"`
	TwDesc  string `json:"twitter_description"`
	TwImage string `json:"twitter_image"`
	TwCard  string `json:"twitter_card"`

	H1 []string `json:"h1"`
	H2 []string `json:"h2"`
	H3 []string `json:"h3"`

	Images   []string `json:"images"`
	Links    []string `json:"links"`
	Favicons []string `json:"favicons"`
	Scripts  []string `json:"scripts"`
}

func (ScrapeResult) TableName() string {
	return "scrape_results"
}
