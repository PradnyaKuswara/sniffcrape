package models

import "gorm.io/gorm"

type ScrapeResult struct {
	gorm.Model
	ID 			uint   `json:"id" gorm:"primaryKey"`
	Url 		string `json:"url"`
	Title 		string `json:"title"`
	Description string `json:"description"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}