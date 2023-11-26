package config

import (
	"os"
)

type DatabaseConfig struct {
	Username     string
	Password     string
	DBName       string
	SSLMode      string
	Host         string
	DatabasePort string
	TimeZone     string
}

func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Username:     os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		Host:         os.Getenv("DB_HOST"),
		DatabasePort: os.Getenv("DB_PORT"),
		SSLMode:      os.Getenv("SSL_MODE"),
		TimeZone:     os.Getenv("TIME_ZONE"),
	}
}
