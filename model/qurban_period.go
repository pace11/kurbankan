package model

import "time"

type QurbanPeriod struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Year        int       `json:"year" gorm:"type:year;not null"`
	StartDate   time.Time `json:"start_date" gorm:"type:date;not null"`
	EndDate     time.Time `json:"end_date" gorm:"type:date;not null"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (QurbanPeriod) TableName() string {
	return "qurban_periods"
}
