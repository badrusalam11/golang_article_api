package main

import (
	"golang_article_api/config" // Updated to match the module name
	"golang_article_api/routes" // Updated to match the module name
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.SetupDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.Run(":1214") // Start the server on port 8080
}
