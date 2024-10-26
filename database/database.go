package database

import (
	"fmt"
	"golang_article_api/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupDatabase initializes the database connection using the configuration
func SetupDatabase() (*gorm.DB, error) {
	// Load the configuration if not already loaded
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
		return nil, err
	}

	// Generate DSN from the loaded config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConfig.Database.Username,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.Server,
		config.AppConfig.Database.Port,
		config.AppConfig.Database.Database, // Update this with the actual database name if it’s dynamic
	)
	fmt.Println("Connecting to database with DSN:", dsn)

	// Open the database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil, err
	}
	return db, nil
}
