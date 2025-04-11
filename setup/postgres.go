package setup

import (
	"fmt"
	"log"

	"github.com/PradnyaKuswara/sniffcrape/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres(cfg *Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password,
		cfg.Database.DBName, cfg.Database.Port, cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.User{})
	return db
}
