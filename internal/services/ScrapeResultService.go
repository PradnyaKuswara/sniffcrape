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

func (s *ScrapeResultService) GetAllScrapeResults() ([]models.ScrapeResult, error) {
	scrapeResults, err := s.ScrapeResultRepositoryInterface.GetAllScrapeResults()
	if err != nil {
		return nil, err
	}
	return scrapeResults, nil
}

func (s *ScrapeResultService) GetScrapeResultByID(id string) (models.ScrapeResult, error) {
	scrapeResult, err := s.ScrapeResultRepositoryInterface.GetScrapeResultByID(id)
	if err != nil {
		return models.ScrapeResult{}, err
	}
	return scrapeResult, nil
}

func (s *ScrapeResultService) CreateScrapeResult(req models.ScrapeResultRequest) (models.ScrapeResult, error) {
	data := models.ScrapeResult{
		Title:       req.Title,
		Url:         req.Url,
		Description: req.Description,
	}

	result, err := s.ScrapeResultRepositoryInterface.CreateScrapeResult(data)
	if err != nil {
		return models.ScrapeResult{}, err
	}

	return result, nil
}
