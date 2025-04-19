package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Get data success",
		"data":    data,
	})
}

func CreatedResponse(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Data created",
		"data":    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":  "error",
		"message": message,
		"data":    nil,
	})
}

func DeleteResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Data deleted",
	})
}

func ValidationErrorResponse(c *gin.Context, errors map[string]string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  "error",
		"message": "Validation error",
		"data":    errors,
	})
}

func PaginatedResponse(c *gin.Context, data any, total int64, page int, limit int) {
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Get data success",
		"data":    data,
		"meta": gin.H{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}
