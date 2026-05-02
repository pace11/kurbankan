package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionItemResponse struct {
	ID            uint       `json:"id"`
	TransactionID uint       `json:"transaction_id"`
	ParticipantID uint       `json:"participant_id"`
	Amount        float64    `json:"amount"`
	PaidAt        *time.Time `json:"paid_at"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type TransactionItem struct {
	ID uint `json:"id" gorm:"primaryKey"`

	TransactionID uint `json:"transaction_id"`
	ParticipantID uint `json:"participant_id"`

	Amount *float64 `json:"amount"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Transaction *Transaction `json:"transaction" gorm:"foreignKey:TransactionID"`
	Participant *Participant `json:"participant" gorm:"foreignKey:ParticipantID"`
}

func (TransactionItem) TableName() string {
	return "transaction_items"
}
