package config

import "os"

type DatabaseConfig struct {
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Username: os.Getenv("USER_NAME"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSL_MODE"),
	}
}
