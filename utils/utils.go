package utils

import (
	"golang_article_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SuccessResponse generates a successful response
func SuccessResponse(c *gin.Context, data interface{}, message string) {
	response := models.Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}

// ErrorResponse generates an error response
func ErrorResponse(c *gin.Context, statusCode int, message string) {
	response := models.Response{
		Status:  "error",
		Message: message,
		Data:    nil,
	}
	c.JSON(statusCode, response)
}
