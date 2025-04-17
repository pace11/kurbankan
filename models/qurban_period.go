package models

import (
	"time"

	"gorm.io/gorm"
)

type QurbanPeriod struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Year        int            `json:"year" gorm:"type:year;not null" binding:"required"`
	StartDate   time.Time      `json:"start_date" gorm:"type:date;not null" binding:"required"`
	EndDate     time.Time      `json:"end_date" gorm:"type:date;not null" binding:"required"`
	Description *string        `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (QurbanPeriod) TableName() string {
	return "qurban_periods"
}
