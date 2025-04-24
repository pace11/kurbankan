package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpResponse(c *gin.Context, data any, code int, entity any, action any, errors map[string]string) {
	response := gin.H{
		"status":  http.StatusText(code),
		"message": StatusMessage(code, action, entity),
	}

	if len(errors) > 0 {
		response["errors"] = errors
		response["data"] = nil
	} else {
		response["data"] = data
	}

	c.JSON(code, response)
}

func StatusMessage(code int, action any, entity any) string {
	switch code {
	case http.StatusOK:
		if action == "get" {
			return fmt.Sprintf("Successfully fetched %s data", entity)
		} else if action == "update" {
			return fmt.Sprintf("Successfully updated %s data", entity)
		}
	case http.StatusBadRequest:
		return "Validation error"
	case http.StatusCreated:
		return fmt.Sprintf("Successfully created %s", entity)
	case http.StatusNoContent:
		return fmt.Sprintf("Successfully deleted %s", entity)
	case http.StatusNotFound:
		return fmt.Sprintf("%s not found", entity)
	case http.StatusInternalServerError:
		return fmt.Sprintf("Internal server error while processing %s", entity)
	}

	return fmt.Sprintf("Action %s for %s returned status %d", action, entity, code)
}

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

func PaginatedResponse(c *gin.Context, data any, code int, entity any, action any, total int64, page int, limit int) {
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(code),
		"message": StatusMessage(code, action, entity),
		"data":    data,
		"meta": gin.H{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}
