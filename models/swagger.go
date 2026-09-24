package models

// Generic swagger response wrappers used across the API.

// PaginatedMeta holds pagination metadata.
type PaginatedMeta struct {
	Page       int   `json:"page" example:"1"`
	Limit      int   `json:"limit" example:"10"`
	Total      int64 `json:"total" example:"100"`
	TotalPages int   `json:"total_pages" example:"10"`
}

// SwaggerError is the standard error object.
type SwaggerError struct {
	Code    string `json:"code" example:"NOT_FOUND"`
	Message string `json:"message" example:"Qurban period not found"`
}

// SwaggerErrorResponse wraps SwaggerError in the API error envelope.
type SwaggerErrorResponse struct {
	Error SwaggerError `json:"error"`
}

// SwaggerValidationDetail holds per-field validation errors.
type SwaggerValidationDetail map[string]string

// SwaggerValidationError is the error object for validation failures.
type SwaggerValidationError struct {
	Code    string                  `json:"code" example:"VALIDATION_ERROR"`
	Message string                  `json:"message" example:"Validation failed"`
	Details SwaggerValidationDetail `json:"details"`
}

// SwaggerValidationErrorResponse wraps SwaggerValidationError.
type SwaggerValidationErrorResponse struct {
	Error SwaggerValidationError `json:"error"`
}
