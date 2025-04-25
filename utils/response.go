package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusMessage(code int, method string, entity any) string {
	messages := map[int]string{
		http.StatusBadRequest:          "Validation error %v",
		http.StatusCreated:             "Successfully created %v",
		http.StatusNotFound:            "%v not found",
		http.StatusInternalServerError: "Internal server error while processing %v",
	}

	methodMessages := map[string]string{
		"GET":    "Successfully fetched %v data",
		"PATCH":  "Successfully updated %v data",
		"DELETE": "Successfully deleted %v data",
	}

	template, ok := messages[code]
	if !ok {
		httpMessage := methodMessages[method]
		template = httpMessage
	}

	return fmt.Sprintf(template, entity)
}

func HttpResponse(c *gin.Context, data any, code int, entity any, method string, errors map[string]string) {
	response := gin.H{
		"status":  http.StatusText(code),
		"message": StatusMessage(code, method, entity),
	}

	if len(errors) > 0 {
		response["errors"] = errors
		response["data"] = nil
	} else {
		response["data"] = data
	}

	c.JSON(code, response)
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

func PaginatedResponse(c *gin.Context, data any, code int, entity any, method string, total int64, page int, limit int) {
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(code),
		"message": StatusMessage(code, method, entity),
		"data":    data,
		"meta": gin.H{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}
