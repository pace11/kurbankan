package models

import (
	"time"

	"gorm.io/gorm"
)

type MosquePaymentType string

const (
	BankTransfer MosquePaymentType = "bank_transfer"
	QRISStatic   MosquePaymentType = "qris_static"
)

type MosquePaymentMethodResponse struct {
	ID            uint              `json:"id"`
	Type          MosquePaymentType `json:"type"`
	Name          string            `json:"name"`
	AccountName   *string           `json:"account_name,omitempty"`
	AccountNumber *string           `json:"account_number,omitempty"`
	QRISImageURL  *string           `json:"qris_image_url,omitempty"`
	Instructions  *string           `json:"instructions,omitempty"`
	IsActive      bool              `json:"is_active"`
	CreatedAt     *time.Time        `json:"created_at"`
	UpdatedAt     *time.Time        `json:"updated_at"`
	Bank          *Bank             `json:"bank,omitempty"`
}

type MosquePaymentMethod struct {
	ID uint `json:"id" gorm:"primaryKey"`

	MosqueID uint  `json:"mosque_id" gorm:"not null"`
	BankID   *uint `json:"bank_id"`

	Type          MosquePaymentType `json:"type" gorm:"type:varchar(15);not null"`
	Name          string            `json:"name" gorm:"size:100;not null"`
	AccountName   *string           `json:"account_name"`
	AccountNumber *string           `json:"account_number"`
	QRISImageURL  *string           `json:"qris_image_url"`
	Instructions  *string           `json:"instructions"`
	IsActive      bool              `json:"is_active" gorm:"default:true"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Mosque *Mosque `json:"mosque" gorm:"foreignKey:MosqueID"`
	Bank   *Bank   `json:"bank" gorm:"foreignKey:BankID"`
}

func (MosquePaymentMethod) TableName() string {
	return "mosque_payment_methods"
}
