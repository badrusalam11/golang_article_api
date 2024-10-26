package controllers

import (
	"fmt"
	"golang_article_api/models" // Updated to match the module name
	"net/http"
	"strconv" // For string-to-int conversions

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateArticle handles creating a new article
func CreateArticle(c *gin.Context, db *gorm.DB) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation
	if len(article.Title) < 20 || len(article.Content) < 200 || len(article.Category) < 3 {
		fmt.Println(len(article.Title))
		fmt.Println(len(article.Content))
		fmt.Println(len(article.Category))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if article.Status != "publish" && article.Status != "draft" && article.Status != "thrash" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	db.Create(&article)
	c.JSON(http.StatusOK, article)
}

// GetArticles handles fetching articles with limit and offset as query parameters
func GetArticles(c *gin.Context, db *gorm.DB) {
	var articles []models.Article

	// Retrieve limit and offset from query parameters
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10")) // Default limit to 10
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0")) // Default offset to 0
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
		return
	}

	// Fetch articles with limit and offset
	db.Limit(limit).Offset(offset).Find(&articles)
	c.JSON(http.StatusOK, articles)
}

// GetArticle fetches a single article by its ID
func GetArticle(c *gin.Context, db *gorm.DB) {
	var article models.Article
	if err := db.First(&article, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	c.JSON(http.StatusOK, article)
}

// UpdateArticle updates an existing article by its ID
func UpdateArticle(c *gin.Context, db *gorm.DB) {
	var article models.Article
	if err := db.First(&article, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	var input models.Article
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation for update
	if len(input.Title) < 20 || len(input.Content) < 200 || len(input.Category) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if input.Status != "publish" && input.Status != "draft" && input.Status != "thrash" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	db.Model(&article).Updates(input)
	c.JSON(http.StatusOK, article)
}

// DeleteArticle deletes an article by its ID
func DeleteArticle(c *gin.Context, db *gorm.DB) {
	var article models.Article
	if err := db.First(&article, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	db.Delete(&article)
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted"})
}
