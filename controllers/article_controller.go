package controllers

import (
	"golang_article_api/models"
	"golang_article_api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateArticle handles creating a new article
func CreateArticle(c *gin.Context, db *gorm.DB) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}

	// Validation
	if len(article.Title) < 20 || len(article.Content) < 200 || len(article.Category) < 3 {
		utils.ErrorResponse(c, http.StatusBadRequest, "title must be at least 20 characters, content 200 characters, and category 3 characters")
		return
	}

	if article.Status != "publish" && article.Status != "draft" && article.Status != "thrash" {
		utils.ErrorResponse(c, http.StatusBadRequest, "status must be publish, draft, or thrash")
		return
	}

	db.Create(&article)
	utils.SuccessResponse(c, article.ToResponse(), "article created successfully")
}

// GetArticles handles fetching articles with limit and offset as query parameters
func GetArticles(c *gin.Context, db *gorm.DB) {
	var articles []models.Article

	// Retrieve limit and offset from query parameters
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid limit parameter")
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid offset parameter")
		return
	}

	// Fetch articles with limit and offset
	db.Limit(limit).Offset(offset).Find(&articles)

	// Convert []Article to []ArticleResponse
	var articleResponses []models.ArticleResponse
	for _, article := range articles {
		articleResponses = append(articleResponses, article.ToResponse())
	}

	// Use SuccessResponse to return standardized response
	utils.SuccessResponse(c, articleResponses, "articles retrieved successfully")
}

// GetArticle fetches a single article by its ID
func GetArticle(c *gin.Context, db *gorm.DB) {
	var article models.Article
	if err := db.First(&article, c.Param("id")).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "article not found")
		return
	}
	utils.SuccessResponse(c, article.ToResponse(), "article retrieved successfully")
}

// UpdateArticle updates an existing article by its ID
func UpdateArticle(c *gin.Context, db *gorm.DB) {
	var article models.Article
	if err := db.First(&article, c.Param("id")).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "article not found")
		return
	}

	var input models.Article
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}

	// Validation for update
	if len(input.Title) < 20 || len(input.Content) < 200 || len(input.Category) < 3 {
		utils.ErrorResponse(c, http.StatusBadRequest, "title must be at least 20 characters, content 200 characters, and category 3 characters")
		return
	}

	if input.Status != "publish" && input.Status != "draft" && input.Status != "thrash" {
		utils.ErrorResponse(c, http.StatusBadRequest, "status must be publish, draft, or thrash")
		return
	}

	db.Model(&article).Updates(input)
	utils.SuccessResponse(c, article, "article updated successfully")
}

// DeleteArticle deletes an article by its ID
func DeleteArticle(c *gin.Context, db *gorm.DB) {
	var article models.Article
	if err := db.First(&article, c.Param("id")).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "article not found")
		return
	}

	db.Delete(&article)
	utils.SuccessResponse(c, nil, "article deleted successfully")
}
