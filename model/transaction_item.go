package model

import "time"

type TransactionStatus string
type PaymentType string

const (
	Pending   TransactionStatus = "pending"
	Paid      TransactionStatus = "paid"
	Cancelled TransactionStatus = "cancelled"

	VA PaymentType = "VA"
)

type TransactionItem struct {
	ID            uint              `json:"id" gorm:"primaryKey"`
	TransactionID uint              `json:"transaction_id"`
	ParticipantID uint              `json:"participant_id"`
	Amount        float64           `json:"amount"`
	Status        TransactionStatus `json:"status" gorm:"type:enum('pending','paid','cancelled');default:'pending'"`
	PaymentType   PaymentType       `json:"payment_type" gorm:"type:enum('VA');default:'VA'"`
	ExternalID    *string           `json:"external_id"`
	PaidAt        *time.Time        `json:"paid_at"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`

	Transaction *Transaction `json:"transaction" gorm:"foreignKey:TransactionID"`
	Participant *Participant `json:"participant" gorm:"foreignKey:ParticipantID"`
}

func (TransactionItem) TableName() string {
	return "transaction_items"
}
