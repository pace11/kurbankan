package model

import "time"

type Beneficiary struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	MosqueID  uint      `json:"mosque_id"`
	Name      string    `json:"name"`
	Address   *string   `json:"address"`
	Phone     *string   `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Mosque *Mosque `json:"mosque" gorm:"foreignKey:MosqueID"`
}

func (Beneficiary) TableName() string {
	return "beneficiaries"
}
