package repositories

import (
	"github.com/PradnyaKuswara/sniffcrape/internal/models"
	"gorm.io/gorm"
)

type ScrapeResultRepository struct {
	DB *gorm.DB
}

type ScrapeResultRepositoryInterface interface {
	GetAllScrapeResults() ([]models.ScrapeResultModel, error)
	GetScrapeResultByID(id string) (models.ScrapeResultModel, error)
	CreateScrapeResult(scrapeResult models.ScrapeResultModel) (models.ScrapeResultModel, error) // ✅ diperbaiki
}

func (r *ScrapeResultRepository) GetAllScrapeResults() ([]models.ScrapeResultModel, error) {
	var scrapeResults []models.ScrapeResultModel
	if err := r.DB.Find(&scrapeResults).Error; err != nil {
		return nil, err
	}
	return scrapeResults, nil
}

func (r *ScrapeResultRepository) GetScrapeResultByID(id string) (models.ScrapeResultModel, error) {
	var scrapeResult models.ScrapeResultModel
	if err := r.DB.First(&scrapeResult, id).Error; err != nil {
		return models.ScrapeResultModel{}, err
	}
	return scrapeResult, nil
}

func (r *ScrapeResultRepository) CreateScrapeResult(data models.ScrapeResultModel) (models.ScrapeResultModel, error) {
	if err := r.DB.Create(&data).Error; err != nil {
		return models.ScrapeResultModel{}, err
	}
	return data, nil
}