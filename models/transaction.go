package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionResponse struct {
	ID             uint      `json:"id"`
	QurbanPeriodID uint      `json:"qurban_period_id"`
	MosqueID       uint      `json:"mosque_id"`
	QurbanOptionID uint      `json:"qurban_option_id"`
	IsFull         bool      `json:"is_full"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	Mosque         *Mosque         `json:"mosque"`
	QurbanPeriod   *QurbanPeriod   `json:"qurban_period"`
	QurbanOffering *QurbanOffering `json:"qurban_offering"`
}

type Transaction struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"size:100;unique"`

	MosqueID         uint `json:"mosque_id"`
	CreatedByUserID  uint `json:"created_by_user_id"`
	QurbanPeriodID   uint `json:"qurban_period_id"`
	QurbanOfferingID uint `json:"qurban_offering_id"`
	IsFull           bool `json:"is_full" gorm:"default:false"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Mosque         *Mosque         `json:"mosque" gorm:"foreignKey:MosqueID"`
	QurbanPeriod   *QurbanPeriod   `json:"qurban_period" gorm:"foreignKey:QurbanPeriodID"`
	QurbanOffering *QurbanOffering `json:"qurban_offering" gorm:"foreignKey:QurbanOfferingID"`
}

func (Transaction) TableName() string {
	return "transactions"
}
