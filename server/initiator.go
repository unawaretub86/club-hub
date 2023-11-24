package server

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/unawaretub86/club-hub/config"
)

func InitDB() (*gorm.DB, error) {
	dbConfig := config.GetDatabaseConfig()

	db, err := gorm.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=%s", dbConfig.Username, dbConfig.DBName, dbConfig.SSLMode))
	if err != nil {
		return nil, err
	}

	return db, nil
}
