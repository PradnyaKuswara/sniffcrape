package setup

import (
	"log"
	"strings"

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

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	overrideWithEnv()

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal("Unable to decode config", err)
	}

	return &AppConfig
}

func overrideWithEnv() {
	if val := viper.GetString("SERVER_PORT"); val != "" {
		viper.Set("server.port", val)
	}
	if val := viper.GetString("DATABASE_HOST"); val != "" {
		viper.Set("database.host", val)
	}
	if val := viper.GetString("DATABASE_PORT"); val != "" {
		viper.Set("database.port", val)
	}
	if val := viper.GetString("DATABASE_USER"); val != "" {
		viper.Set("database.user", val)
	}
	if val := viper.GetString("DATABASE_PASSWORD"); val != "" {
		viper.Set("database.password", val)
	}
	if val := viper.GetString("DATABASE_DBNAME"); val != "" {
		viper.Set("database.dbname", val)
	}
	if val := viper.GetString("DATABASE_SSLMODE"); val != "" {
		viper.Set("database.sslmode", val)
	}
}
