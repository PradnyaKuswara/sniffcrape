package routes

import (
	"github.com/PradnyaKuswara/sniffcrape/internal/handlers"
	"github.com/PradnyaKuswara/sniffcrape/internal/middlewares"
	"github.com/PradnyaKuswara/sniffcrape/internal/repositories"
	"github.com/PradnyaKuswara/sniffcrape/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterScrapeResult(r *gin.Engine, db *gorm.DB) {
	repo := &repositories.ScrapeResultRepository{DB: db}
	scrapeResultService := services.NewScrapeResultService(repo)
	scrapeResultHandler := handlers.NewScrapeResultHandler(scrapeResultService)

	scrapeResultRoutes := r.Group("/api/v1/scrape-results").Use(middlewares.AuthMiddleware())
	{
		scrapeResultRoutes.POST("/colly-scrap", func(c *gin.Context) {
			scrapeResultHandler.CollyScrap(c)
		})

	}
}
