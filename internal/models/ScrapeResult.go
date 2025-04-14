package models

import "gorm.io/gorm"

// type ScrapeResult struct {
// 	ID          uint   `json:"id"`
// 	Url         string `json:"url"`
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// 	CreatedAt   string `json:"created_at"`
// 	UpdatedAt   string `json:"updated_at"`
// 	DeletedAt   string `json:"deleted_at,omitempty"`
// }

type ScrapeResultModel struct {
	gorm.Model
	Url 		string `json:"url"`
	Title 		string `json:"title"`
	Description string `json:"description"`
}

type ScrapeResultRequest struct {
	Url         string `json:"url" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (ScrapeResultModel) TableName() string {
	return "scrape_results"
}

