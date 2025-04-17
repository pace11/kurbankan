package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	QurbanPeriodID uint           `json:"qurban_period_id"`
	MosqueID       uint           `json:"mosque_id"`
	QurbanOptionID uint           `json:"qurban_option_id"`
	IsFull         bool           `json:"is_full" gorm:"default:false"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Mosque       *Mosque       `json:"mosque" gorm:"foreignKey:MosqueID"`
	QurbanPeriod *QurbanPeriod `json:"qurban_period" gorm:"foreignKey:QurbanPeriodID"`
	QurbanOption *QurbanOption `json:"qurban_option" gorm:"foreignKey:QurbanOptionID"`
}

func (Transaction) TableName() string {
	return "transactions"
}
