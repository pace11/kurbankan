package models

import (
	"time"

	"gorm.io/gorm"
)

type QurbanAnimalType string
type QurbanSchemeType string
type QurbanStatus string

const (
	Cow  QurbanAnimalType = "cow"
	Goat QurbanAnimalType = "goat"

	Group      QurbanSchemeType = "group"
	Individual QurbanSchemeType = "individual"

	Open   QurbanStatus = "open"
	Closed QurbanStatus = "closed"
)

// ========== Models ==========

type QurbanOffering struct {
	ID uint `json:"id" gorm:"primaryKey"`

	MosqueID       uint `json:"mosque_id"`
	QurbanPeriodID uint `json:"qurban_period_id" binding:"required"`

	AnimalType     QurbanAnimalType `json:"animal_type" gorm:"type:varchar(20);not null" binding:"required"`
	SchemeType     QurbanSchemeType `json:"scheme_type" gorm:"type:varchar(20);not null" binding:"required"`
	Name           string           `json:"name" gorm:"size:100" binding:"required"`
	Price          float64          `json:"price" binding:"required"`
	Capacity       int              `json:"capacity" binding:"required"`
	FilledSlots    int              `json:"filled_slots" gorm:"default:0"`
	ConfirmedSlots int              `json:"confirmed_slots" gorm:"default:0"`
	Status         QurbanStatus     `json:"status" gorm:"type:varchar(20);default:'open'"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	QurbanPeriod *QurbanPeriod `json:"qurban_period" gorm:"foreignKey:QurbanPeriodID"`
	Mosque       *Mosque       `json:"mosque" gorm:"foreignKey:MosqueID"`
}

func (QurbanOffering) TableName() string {
	return "qurban_offerings"
}

// ========== Responses (Response DTOs) ==========

// QurbanOfferingResponse is the single-item response for qurban offerings.
type QurbanOfferingResponse struct {
	ID             uint             `json:"id"`
	QurbanPeriodID uint             `json:"qurban_period_id"`
	AnimalType     QurbanAnimalType `json:"animal_type"`
	SchemeType     QurbanSchemeType `json:"scheme_type"`
	Name           string           `json:"name"`
	Price          float64          `json:"price"`
	Capacity       int              `json:"capacity"`
	FilledSlots    int              `json:"filled_slots"`
	ConfirmedSlots int              `json:"confirmed_slots"`
	Status         QurbanStatus     `json:"status"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
}

// ========== Payloads (Request DTOs) ==========

type QurbanOfferingRequest struct {
	QurbanPeriodID uint             `json:"qurban_period_id" binding:"required"`
	AnimalType     QurbanAnimalType `json:"animal_type" binding:"required,oneof=cow goat"`
	SchemeType     QurbanSchemeType `json:"scheme_type" binding:"required,oneof=group individual"`
	Name           string           `json:"name" binding:"required,max=100"`
	Price          float64          `json:"price" binding:"required,gt=0"`
	Capacity       int              `json:"capacity" binding:"required,gt=0"`
	Status         QurbanStatus     `json:"status" binding:"omitempty,oneof=open closed"`

	// Not provided in payload
	MosqueID uint `json:"-"`
	ID       uint `json:"-"`
}

// QurbanOfferingListWithPaginationResponse is the paginated list response for qurban offerings.
type QurbanOfferingListWithPaginationResponse struct {
	Data []QurbanOfferingResponse `json:"data"`
	Meta PaginatedMeta            `json:"meta"`
}

// QurbanOfferingDetailResponse is the single-item GET response.
type QurbanOfferingDetailResponse struct {
	Data QurbanOfferingResponse `json:"data"`
}

// QurbanOfferingMutationResponse is the create/update/delete response.
type QurbanOfferingMutationResponse struct {
	Message string                 `json:"message" example:"Qurban Offering created successfully"`
	Data    QurbanOfferingResponse `json:"data"`
}
