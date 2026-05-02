package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Compile regex once at package level for better performance
var fieldExtractorRegex = regexp.MustCompile(`field \w+\.(\w+)`)

// getJSONFieldName extracts JSON field name from struct field using reflection
// Returns the JSON tag name if available, otherwise returns the original field name
func getJSONFieldName(refType reflect.Type, fieldName string) string {
	field, ok := refType.FieldByName(fieldName)
	if !ok {
		return fieldName
	}

	jsonTag := field.Tag.Get("json")
	if jsonTag == "" || jsonTag == "-" {
		return fieldName
	}

	// Split by comma to handle tags like `json:"name,omitempty"`
	parts := strings.Split(jsonTag, ",")
	if parts[0] != "" {
		return parts[0]
	}

	return fieldName
}

func BindAndValidate(c *gin.Context, form any) error {
	if err := c.ShouldBindJSON(form); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make(map[string]string, len(ve))
			refType := reflect.TypeOf(form).Elem()

			for _, fe := range ve {
				jsonField := getJSONFieldName(refType, fe.Field())
				out[jsonField] = getErrorMessage(fe, jsonField)
			}

			HttpResponse(c, nil, http.StatusBadRequest, "data", c.Request.Method, out)
			return err
		}

		// Handle JSON unmarshal type errors
		var unmarshalTypeError *json.UnmarshalTypeError
		if errors.As(err, &unmarshalTypeError) {
			errorMap := parseUnmarshalError(unmarshalTypeError, reflect.TypeOf(form).Elem())
			HttpResponse(c, nil, http.StatusBadRequest, "data", c.Request.Method, errorMap)
			return err
		}

		// Handle other JSON syntax errors
		errorMap := parseGenericJSONError(err.Error())
		HttpResponse(c, nil, http.StatusBadRequest, "data", c.Request.Method, errorMap)
		return err
	}
	return nil
}

// parseUnmarshalError converts JSON unmarshal type error to user-friendly message
func parseUnmarshalError(err *json.UnmarshalTypeError, refType reflect.Type) map[string]string {
	jsonField := getJSONFieldName(refType, err.Field)
	expectedType := getTypeName(err.Type.String())

	return map[string]string{
		jsonField: fmt.Sprintf("%s must be a %s, got %s", jsonField, expectedType, err.Value),
	}
}

// parseGenericJSONError parses generic JSON errors and extracts field names
func parseGenericJSONError(errMsg string) map[string]string {
	// Try to extract field name from error message
	// Pattern: "json: cannot unmarshal X into Go struct field Y.Z of type T"
	matches := fieldExtractorRegex.FindStringSubmatch(errMsg)

	if len(matches) > 1 {
		fieldName := toSnakeCase(matches[1])
		return map[string]string{
			fieldName: fmt.Sprintf("%s has invalid type or format", fieldName),
		}
	}

	// Fallback to generic error
	return map[string]string{
		"error": "Invalid request format",
	}
}

// getTypeName converts Go type name to user-friendly name
func getTypeName(goType string) string {
	switch {
	case strings.Contains(goType, "uint"):
		return "number"
	case strings.Contains(goType, "int"):
		return "number"
	case strings.Contains(goType, "float"):
		return "number"
	case strings.Contains(goType, "bool"):
		return "boolean"
	case strings.Contains(goType, "string"):
		return "string"
	default:
		return "valid value"
	}
}

// toSnakeCase converts camelCase or PascalCase to snake_case
// Handles consecutive uppercase letters correctly (e.g., "HTTPServer" -> "http_server")
func toSnakeCase(str string) string {
	if str == "" {
		return ""
	}

	// Pre-allocate with estimated capacity (original length + 30% for underscores)
	var result strings.Builder
	result.Grow(len(str) + len(str)/3)

	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		r := runes[i]

		// Add underscore before uppercase letter if:
		// 1. Not the first character
		// 2. Previous character is lowercase OR
		// 3. Next character exists and is lowercase (handles "HTTPServer" -> "http_server")
		if i > 0 && unicode.IsUpper(r) {
			prevIsLower := unicode.IsLower(runes[i-1])
			nextIsLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])

			if prevIsLower || nextIsLower {
				result.WriteRune('_')
			}
		}

		result.WriteRune(unicode.ToLower(r))
	}

	return result.String()
}

// getErrorMessage returns a user-friendly error message based on validation tag
func getErrorMessage(fe validator.FieldError, jsonField string) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", jsonField)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", jsonField, fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", jsonField, fe.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", jsonField)
	case "oneof":
		return fmt.Sprintf("%s must be one of [%s]", jsonField, fe.Param())
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", jsonField, fe.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", jsonField, fe.Param())
	case "lt":
		return fmt.Sprintf("%s must be less than %s", jsonField, fe.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", jsonField, fe.Param())
	case "url":
		return fmt.Sprintf("%s must be a valid URL", jsonField)
	case "ne":
		return fmt.Sprintf("%s cannot be empty", jsonField)
	default:
		return fmt.Sprintf("%s is invalid", jsonField)
	}
}
