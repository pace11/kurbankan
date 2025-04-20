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
