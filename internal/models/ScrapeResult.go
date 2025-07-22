package models

import (
	"time"

	"gorm.io/gorm"
)

// ScrapeResultModel digunakan untuk operasi ke database
type ScrapeResult struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Url         string         `json:"url"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// ScrapeResultRequest digunakan untuk validasi input dari client
type ScrapeResultRequest struct {
	Url         string `json:"url" validate:"required"` // pakai validator, bukan binding (sesuai obrolan sebelumnya)
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// TableName memastikan nama tabel sesuai
func (ScrapeResult) TableName() string {
	return "scrape_results"
}
