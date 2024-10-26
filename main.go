package main

import (
	// Updated to match the module name
	"fmt"
	"golang_article_api/config"
	"golang_article_api/database"
	"golang_article_api/routes" // Updated to match the module name
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load the configuration
	err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	db, err := database.SetupDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.Run(":1214") // Start the server on port 8080
}
