package main

import (
	"github.com/PradnyaKuswara/sniffcrape/internal/setup"
	"github.com/PradnyaKuswara/sniffcrape/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	logger.InitLogger()

	err := godotenv.Load()
	if err != nil {
		logger.Log.Warn("Error loading .env file")
	} else {
		logger.Log.Info(".env file loaded successfully")
	}

	cfg := setup.InitConfig()
	db := setup.InitPostgres(cfg)
	router := setup.InitRouter(db)

	logger.Log.Infof("Starting server on port: %s", cfg.Server.Port)
	router.Run(":" + cfg.Server.Port)
}
