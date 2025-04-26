package models

import (
	"time"
)

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type MosqueResponse struct {
	ID        uint          `json:"id"`
	Name      string        `json:"name"`
	Address   *string       `json:"address"`
	Photos    *string       `json:"photos"`
	Province  *Province     `json:"province"`
	Regency   *Regency      `json:"regency"`
	District  *District     `json:"district"`
	Village   *Village      `json:"village"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	User      *UserResponse `json:"user"`
}
type ParticipantResponse struct {
	ID        uint          `json:"id"`
	Name      string        `json:"name"`
	Address   *string       `json:"address"`
	Province  *Province     `json:"province"`
	Regency   *Regency      `json:"regency"`
	District  *District     `json:"district"`
	Village   *Village      `json:"village"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	User      *UserResponse `json:"user"`
}

type BeneficiaryResponse struct {
	ID        uint      `json:"id"`
	MosqueID  uint      `json:"mosque_id"`
	Name      string    `json:"name"`
	Address   *string   `json:"address"`
	Phone     *string   `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TransactionResponse struct {
	ID             uint      `json:"id"`
	QurbanPeriodID uint      `json:"qurban_period_id"`
	MosqueID       uint      `json:"mosque_id"`
	QurbanOptionID uint      `json:"qurban_option_id"`
	IsFull         bool      `json:"is_full"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	Mosque       *Mosque       `json:"mosque"`
	QurbanPeriod *QurbanPeriod `json:"qurban_period"`
	QurbanOption *QurbanOption `json:"qurban_option"`
}
type TransactionItemResponse struct {
	ID            uint              `json:"id"`
	TransactionID uint              `json:"transaction_id"`
	ParticipantID uint              `json:"participant_id"`
	Amount        float64           `json:"amount"`
	Status        TransactionStatus `json:"status"`
	PaymentType   PaymentType       `json:"payment_type"`
	ExternalID    string            `json:"external_id"`
	PaidAt        *time.Time        `json:"paid_at"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
}

type QurbanPeriodResponse struct {
	ID          uint      `json:"id"`
	Year        int       `json:"year"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type QurbanOptionResponse struct {
	ID             uint             `json:"id"`
	QurbanPeriodID uint             `json:"qurban_period_id"`
	AnimalType     QurbanAnimalType `json:"animal_type"`
	SchemeType     QurbanSchemeType `json:"scheme_type"`
	Price          float64          `json:"price"`
	Slots          int              `json:"slots"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
}
type UserListResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
