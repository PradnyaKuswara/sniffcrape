package setup

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
}

var AppConfig Config

func InitConfig() *Config {
	viper.SetConfigFile("configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading config file", err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal("Unable to decode config", err)
	}

	return &AppConfig
}
