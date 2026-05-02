package models

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

// ========== Constants ==========

type TransactionPaymentStatus string

const (
	TransactionPending             TransactionPaymentStatus = "pending"
	TransactionWaitingVerification TransactionPaymentStatus = "waiting_verification"
	TransactionPaid                TransactionPaymentStatus = "paid"
	TransactionRejected            TransactionPaymentStatus = "rejected"
	TransactionCancelled           TransactionPaymentStatus = "cancelled"
)

// ========== Models ==========

type Transaction struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"size:100;unique"`

	MosqueID              uint  `json:"mosque_id"`
	CreatedByUserID       uint  `json:"created_by_user_id"`
	QurbanPeriodID        uint  `json:"qurban_period_id"`
	QurbanOfferingID      uint  `json:"qurban_offering_id"`
	MosquePaymentMethodID *uint `json:"mosque_payment_method_id"`

	PaymentStatus    TransactionPaymentStatus `json:"payment_status" gorm:"type:varchar(30);default:'pending'"`
	ProofURL         *string                  `json:"proof_url"`
	PaidAmount       *float64                 `json:"paid_amount"`
	PaidAt           *time.Time               `json:"paid_at"`
	PaymentNote      *string                  `json:"payment_note"`
	VerifiedByUserID *uint                    `json:"verified_by_user_id"`
	VerifiedAt       *time.Time               `json:"verified_at"`
	RejectedReason   *string                  `json:"rejected_reason"`
	ExternalID       *string                  `json:"external_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Mosque              *Mosque              `json:"mosque" gorm:"foreignKey:MosqueID"`
	QurbanPeriod        *QurbanPeriod        `json:"qurban_period" gorm:"foreignKey:QurbanPeriodID"`
	QurbanOffering      *QurbanOffering      `json:"qurban_offering" gorm:"foreignKey:QurbanOfferingID"`
	MosquePaymentMethod *MosquePaymentMethod `json:"mosque_payment_method" gorm:"foreignKey:MosquePaymentMethodID"`
}

func (Transaction) TableName() string {
	return "transactions"
}

// BeforeCreate hook to auto-generate transaction code
func (t *Transaction) BeforeCreate(tx *gorm.DB) error {
	// Only generate code if not already set
	if t.Code != "" {
		return nil
	}

	// Fetch QurbanOffering to get AnimalType and SchemeType
	var offering QurbanOffering
	if err := tx.First(&offering, t.QurbanOfferingID).Error; err != nil {
		return err
	}

	// Generate transaction code: TRX-{ANIMAL_TYPE}-{SCHEME_TYPE}-{TIMESTAMP}
	timestamp := time.Now().Format("20060102150405")
	animalType := strings.ToUpper(string(offering.AnimalType))
	schemeType := strings.ToUpper(string(offering.SchemeType))

	t.Code = fmt.Sprintf("TRX-%s-%s-%s", animalType, schemeType, timestamp)

	return nil
}

// ========== Payloads (Request DTOs) ==========

// ParticipantQuickCreatePayload for creating participant on-the-fly during transaction
type ParticipantQuickCreatePayload struct {
	Name          string  `json:"name" binding:"required,min=1,max=100"`
	Gender        *Gender `json:"gender" binding:"omitempty,oneof=male female"`
	IsLikeAddress bool    `json:"is_like_address"`
}

// TransactionCreatePayload for creating new transaction
type TransactionCreatePayload struct {
	MosqueID              uint                            `json:"mosque_id" binding:"required"`
	QurbanPeriodID        uint                            `json:"qurban_period_id" binding:"required"`
	QurbanOfferingID      uint                            `json:"qurban_offering_id" binding:"required"`
	MosquePaymentMethodID *uint                           `json:"mosque_payment_method_id"`
	ParticipantIDs        []uint                          `json:"participant_ids"`
	Participants          []ParticipantQuickCreatePayload `json:"participants"`
	PaymentNote           *string                         `json:"payment_note"`

	// Not provided in payload
	CreatedByUserID uint `json:"-"`
}

// TransactionUploadProofPayload for uploading payment proof
type TransactionUploadProofPayload struct {
	ProofURL    string  `json:"proof_url" binding:"required,url"`
	PaidAmount  float64 `json:"paid_amount" binding:"required,gt=0"`
	PaymentNote *string `json:"payment_note"`
}

// TransactionVerifyPayload for verifying payment (mosque/admin)
type TransactionVerifyPayload struct {
	PaymentStatus  TransactionPaymentStatus `json:"payment_status" binding:"required,oneof=paid rejected"`
	RejectedReason *string                  `json:"rejected_reason"` // Required if status is rejected
}

// Legacy: Keep for backward compatibility
type TransactionPayload struct {
	MosqueID         uint    `json:"mosque_id" binding:"required"`
	QurbanOfferingID uint    `json:"qurban_offering_id" binding:"required"`
	ParticipantID    uint    `json:"participant_id" binding:"required"`
	Amount           float64 `json:"amount" binding:"required"`
}

// ========== Responses (Response DTOs) ==========

type TransactionResponse struct {
	ID               uint      `json:"id"`
	Code             string    `json:"code"`
	QurbanPeriodID   uint      `json:"qurban_period_id"`
	MosqueID         uint      `json:"mosque_id"`
	QurbanOfferingID uint      `json:"qurban_offering_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	Mosque         *Mosque         `json:"mosque"`
	QurbanPeriod   *QurbanPeriod   `json:"qurban_period"`
	QurbanOffering *QurbanOffering `json:"qurban_offering"`
}
