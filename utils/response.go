package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// PaginatedResponse for list endpoints: { "data": [], "meta": {...} }
func PaginatedResponse(c *gin.Context, data any, total int64, page int, limit int) {
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"meta": gin.H{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}

// DetailResponse for single GET: { "data": {} }
func DetailResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// MutationResponse for create/update/delete: { "message": "...", "data": {} }
func MutationResponse(c *gin.Context, httpCode int, message string, data any) {
	c.JSON(httpCode, gin.H{
		"message": message,
		"data":    data,
	})
}

// ErrorResponse for single error: { "error": { "code": "...", "message": "..." } }
func ErrorResponse(c *gin.Context, httpCode int, errorCode string, message string) {
	c.JSON(httpCode, gin.H{
		"error": gin.H{
			"code":    errorCode,
			"message": message,
		},
	})
}

// ValidationErrorResponse for validation errors with field details:
// { "error": { "code": "VALIDATION_ERROR", "message": "Validation failed", "details": {...} } }
func ValidationErrorResponse(c *gin.Context, details map[string]string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": gin.H{
			"code":    "VALIDATION_ERROR",
			"message": "Validation failed",
			"details": details,
		},
	})
}

// HandleRepoError handles an error response from a repository call.
// Returns true if an error was detected and the response was written.
func HandleRepoError(c *gin.Context, httpCode int, errors map[string]string) bool {
	if httpCode >= 200 && httpCode < 300 {
		return false
	}
	errorCode := HTTPCodeToErrorCode(httpCode)
	if len(errors) > 0 {
		if msg, ok := errors["error"]; ok && len(errors) == 1 {
			ErrorResponse(c, httpCode, errorCode, msg)
		} else {
			ValidationErrorResponse(c, errors)
		}
	} else {
		ErrorResponse(c, httpCode, errorCode, http.StatusText(httpCode))
	}
	return true
}

// MutationMessage generates a human-readable action message from entity and HTTP method.
func MutationMessage(entity, method string) string {
	name := entityToTitle(entity)
	switch method {
	case "POST":
		return name + " created successfully"
	case "PUT", "PATCH":
		return name + " updated successfully"
	case "DELETE":
		return name + " deleted successfully"
	default:
		return name + " processed successfully"
	}
}

// HTTPCodeToErrorCode maps an HTTP status code to an error code string.
func HTTPCodeToErrorCode(code int) string {
	switch code {
	case http.StatusBadRequest:
		return "BAD_REQUEST"
	case http.StatusUnauthorized:
		return "UNAUTHORIZED"
	case http.StatusForbidden:
		return "FORBIDDEN"
	case http.StatusNotFound:
		return "NOT_FOUND"
	case http.StatusConflict:
		return "CONFLICT"
	case http.StatusUnprocessableEntity:
		return "UNPROCESSABLE_ENTITY"
	case http.StatusInternalServerError:
		return "INTERNAL_SERVER_ERROR"
	default:
		return "ERROR"
	}
}

func entityToTitle(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	words := strings.Fields(s)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
		}
	}
	return strings.Join(words, " ")
}
