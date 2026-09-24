package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MosqueResponse struct {
	ID        uint          `json:"id"`
	UUID      string        `json:"uuid"`
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

type Mosque struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	UUID   string `json:"uuid" gorm:"type:char(36);uniqueIndex;not null"`
	UserID uint   `json:"user_id"`

	Name         string  `json:"name" gorm:"type:text;not null"`
	Address      *string `json:"address"`
	Photos       *string `json:"photos"`
	ProvinceCode string  `json:"province_code"`
	RegencyCode  string  `json:"regency_code"`
	DistrictCode string  `json:"district_code"`
	VillageCode  string  `json:"village_code"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	User     *User     `json:"user" gorm:"foreignKey:UserID"`
	Province *Province `json:"province" gorm:"foreignKey:ProvinceCode;references:Code"`
	Regency  *Regency  `json:"regency" gorm:"foreignKey:RegencyCode;references:Code"`
	District *District `json:"district" gorm:"foreignKey:DistrictCode;references:Code"`
	Village  *Village  `json:"village" gorm:"foreignKey:VillageCode;references:Code"`
}

func (Mosque) TableName() string {
	return "mosques"
}

func (m *Mosque) BeforeCreate(tx *gorm.DB) error {
	m.UUID = uuid.New().String()
	return nil
}
