package routes

import (
	"golang_article_api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Create a new article
	r.POST("/article", func(c *gin.Context) {
		controllers.CreateArticle(c, db)
	})

	// Get articles with limit and offset as query parameters
	r.GET("/articles", func(c *gin.Context) {
		controllers.GetArticles(c, db)
	})

	// Get a single article by ID
	r.GET("/article/:id", func(c *gin.Context) {
		controllers.GetArticle(c, db)
	})

	// Update an article by ID
	r.PUT("/article/:id", func(c *gin.Context) {
		controllers.UpdateArticle(c, db)
	})

	// Delete an article by ID
	r.DELETE("/article/:id", func(c *gin.Context) {
		controllers.DeleteArticle(c, db)
	})
}
