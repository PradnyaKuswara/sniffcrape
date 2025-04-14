package main

import (
	"log"

	"github.com/PradnyaKuswara/sniffcrape/setup"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file")
    }
	log.Println("Load .env file")
	cfg := setup.InitConfig()
	db := setup.InitPostgres(cfg)
	router := setup.InitRouter(db)

	// Start the server
	router.Run(":" + cfg.Server.Port)
}