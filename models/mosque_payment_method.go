package models

import (
	"time"

	"gorm.io/gorm"
)

// ========== Constants ==========

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

// ========== Models ==========

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

// ========== Payloads (Request DTOs) ==========

type MosquePaymentMethodCreatePayload struct {
	Type          MosquePaymentType `json:"type" binding:"required,oneof=bank_transfer qris_static"`
	Name          string            `json:"name" binding:"required,min=1,max=100"`
	AccountName   *string           `json:"account_name"`   // Required if type is bank_transfer
	AccountNumber *string           `json:"account_number"` // Required if type is bank_transfer
	QRISImageURL  *string           `json:"qris_image_url"` // Required if type is qris_static
	Instructions  *string           `json:"instructions"`
	IsActive      *bool             `json:"is_active"`
	BankID        *uint             `json:"bank_id"` // Optional: only for bank_transfer

	// Not provided in payload
	MosqueID uint `json:"-"`
}
