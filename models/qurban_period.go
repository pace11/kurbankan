package models

import (
	"time"

	"gorm.io/gorm"
)

// ========== Models ==========

type QurbanPeriod struct {
	ID uint `json:"id" gorm:"primaryKey"`

	MosqueID uint `json:"mosque_id"`

	Year        int       `json:"year" gorm:"type:integer;not null" binding:"required"`
	StartDate   time.Time `json:"start_date" gorm:"type:date;not null" binding:"required"`
	EndDate     time.Time `json:"end_date" gorm:"type:date;not null" binding:"required"`
	Description *string   `json:"description"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Mosque *Mosque `json:"mosque" gorm:"foreignKey:MosqueID"`
}

func (QurbanPeriod) TableName() string {
	return "qurban_periods"
}

// ========== Payloads (Request DTOs) ==========

type QurbanPeriodRequest struct {
	Year        int       `json:"year" binding:"required"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
	Description *string   `json:"description,omitempty"`

	// Not provided in payload
	MosqueID uint `json:"-"`
	ID       uint `json:"-"`
}

// ========== Responses (Response DTOs) ==========

// QurbanPeriodResponse is the single-item response for qurban periods.
type QurbanPeriodResponse struct {
	ID          uint      `json:"id"`
	Year        int       `json:"year"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// QurbanPeriodOptionsResponse is the response for qurban period options.
type QurbanPeriodOptionsResponse struct {
	Data []QurbanPeriodResponse `json:"data"`
}

// QurbanPeriodListWithPaginationResponse is the paginated list response for qurban periods.
type QurbanPeriodListWithPaginationResponse struct {
	Data []QurbanPeriodResponse `json:"data"`
	Meta PaginatedMeta          `json:"meta"`
}

// QurbanPeriodDetailResponse is the single-item GET response.
type QurbanPeriodDetailResponse struct {
	Data QurbanPeriodResponse `json:"data"`
}

// QurbanPeriodMutationResponse is the create/update/delete response.
type QurbanPeriodMutationResponse struct {
	Message string               `json:"message" example:"Qurban Period created successfully"`
	Data    QurbanPeriodResponse `json:"data"`
}
