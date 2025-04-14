package handlers

import (
	"net/http"

	"github.com/PradnyaKuswara/sniffcrape/internal/models"
	"github.com/PradnyaKuswara/sniffcrape/internal/services"
	"github.com/PradnyaKuswara/sniffcrape/pkg/utils"
	"github.com/gin-gonic/gin"
)

type ScrapeResultHandler struct {
	ScrapeResultService *services.ScrapeResultService
}

func NewScrapeResultHandler(scrapeResultService *services.ScrapeResultService) *ScrapeResultHandler {
	return &ScrapeResultHandler{
		ScrapeResultService: scrapeResultService,
	}
}

func (s *ScrapeResultHandler) GetAllScrapeResults(c *gin.Context) {
	scrapResults, err := s.ScrapeResultService.GetAllScrapeResults()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get scrap results")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, scrapResults)
}

func (s *ScrapeResultHandler) CreateScrapeResult(c *gin.Context) {
	var scrapResult models.ScrapeResultRequest

	if err := c.ShouldBindJSON(&scrapResult); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	result, err := s.ScrapeResultService.CreateScrapeResult(scrapResult)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create scrap result")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, result)
}
