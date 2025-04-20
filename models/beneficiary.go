package models

import (
	"time"

	"gorm.io/gorm"
)

type Beneficiary struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	MosqueID  uint           `json:"mosque_id" binding:"required"`
	Name      string         `json:"name" binding:"required"`
	Address   *string        `json:"address" binding:"required"`
	Phone     *string        `json:"phone" binding:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Mosque *Mosque `json:"mosque" gorm:"foreignKey:MosqueID"`
}

func (Beneficiary) TableName() string {
	return "beneficiaries"
}
