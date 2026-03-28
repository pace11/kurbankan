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

type QurbanOfferingResponse struct {
	ID             uint             `json:"id"`
	QurbanPeriodID uint             `json:"qurban_period_id"`
	AnimalType     QurbanAnimalType `json:"animal_type"`
	SchemeType     QurbanSchemeType `json:"scheme_type"`
	Name           string           `json:"name"`
	Price          float64          `json:"price"`
	Capacity       int              `json:"capacity"`
	FilledSlots    int              `json:"filled_slots"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
}

type QurbanOffering struct {
	ID uint `json:"id" gorm:"primaryKey"`

	MosqueID       uint `json:"mosque_id"`
	QurbanPeriodID uint `json:"qurban_period_id" binding:"required"`

	AnimalType  QurbanAnimalType `json:"animal_type" gorm:"type:varchar(20);not null" binding:"required"`
	SchemeType  QurbanSchemeType `json:"scheme_type" gorm:"type:varchar(20);not null" binding:"required"`
	Name        string           `json:"name" gorm:"size:100" binding:"required"`
	Price       float64          `json:"price" binding:"required"`
	Capacity    int              `json:"capacity" binding:"required"`
	FilledSlots int              `json:"filled_slots" gorm:"default:0"`
	Status      QurbanStatus     `json:"status" gorm:"type:varchar(20);default:'open'"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	QurbanPeriod *QurbanPeriod `json:"qurban_period" gorm:"foreignKey:QurbanPeriodID"`
	Mosque       *Mosque       `json:"mosque" gorm:"foreignKey:MosqueID"`
}

func (QurbanOffering) TableName() string {
	return "qurban_offerings"
}
