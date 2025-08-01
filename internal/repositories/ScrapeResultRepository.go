package repositories

import (
	"github.com/PradnyaKuswara/sniffcrape/internal/models"
	"gorm.io/gorm"
)

type ScrapeResultRepository struct {
	DB *gorm.DB
}

type ScrapeResultRepositoryInterface interface {
	GetAllScrapeResults() ([]models.ScrapeResult, error)
	GetScrapeResultByID(id string) (models.ScrapeResult, error)
	CreateScrapeResult(scrapeResult models.ScrapeResult) (models.ScrapeResult, error) // ✅ diperbaiki
}

func (r *ScrapeResultRepository) GetAllScrapeResults() ([]models.ScrapeResult, error) {
	var scrapeResults []models.ScrapeResult
	if err := r.DB.Find(&scrapeResults).Error; err != nil {
		return nil, err
	}
	return scrapeResults, nil
}

func (r *ScrapeResultRepository) GetScrapeResultByID(id string) (models.ScrapeResult, error) {
	var scrapeResult models.ScrapeResult
	if err := r.DB.First(&scrapeResult, id).Error; err != nil {
		return models.ScrapeResult{}, err
	}
	return scrapeResult, nil
}

func (r *ScrapeResultRepository) CreateScrapeResult(data models.ScrapeResult) (models.ScrapeResult, error) {
	if err := r.DB.Create(&data).Error; err != nil {
		return models.ScrapeResult{}, err
	}
	return data, nil
}
