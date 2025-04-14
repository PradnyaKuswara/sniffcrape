package services

import (
	"github.com/PradnyaKuswara/sniffcrape/internal/models"
	"github.com/PradnyaKuswara/sniffcrape/internal/repositories"
)

type ScrapeResultService struct {
	ScrapeResultRepositoryInterface repositories.ScrapeResultRepositoryInterface
}

func NewScrapeResultService(repo repositories.ScrapeResultRepositoryInterface) *ScrapeResultService {
	return &ScrapeResultService{
		ScrapeResultRepositoryInterface: repo,
	}
}

func (s *ScrapeResultService) GetAllScrapeResults() ([]models.ScrapeResultModel, error) {
	scrapeResults, err := s.ScrapeResultRepositoryInterface.GetAllScrapeResults()
	if err != nil {
		return nil, err
	}
	return scrapeResults, nil
}

func (s *ScrapeResultService) GetScrapeResultByID(id string) (models.ScrapeResultModel, error) {
	scrapeResult, err := s.ScrapeResultRepositoryInterface.GetScrapeResultByID(id)
	if err != nil {
		return models.ScrapeResultModel{}, err
	}
	return scrapeResult, nil
}

func (s *ScrapeResultService) CreateScrapeResult(req models.ScrapeResultRequest) (models.ScrapeResultModel, error) {
	data := models.ScrapeResultModel{
		Title: req.Title,
		Url:   req.Url,
		Description: req.Description,
	}

	result, err := s.ScrapeResultRepositoryInterface.CreateScrapeResult(data)
	if err != nil {
		return models.ScrapeResultModel{}, err
	}

	return result, nil
}
