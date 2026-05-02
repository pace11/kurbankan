package models

import (
	"time"

	"gorm.io/gorm"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

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

type Participant struct {
	ID              uint  `json:"id" gorm:"primaryKey"`
	UserID          *uint `json:"user_id"`
	CreatedByUserID uint  `json:"created_by_user_id"`

	Name         string  `json:"name"`
	Address      *string `json:"address"`
	Gender       *Gender `json:"gender" gorm:"type:varchar(10)"`
	ProvinceCode *string `json:"province_code"`
	RegencyCode  *string `json:"regency_code"`
	DistrictCode *string `json:"district_code"`
	VillageCode  *string `json:"village_code"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	User     *User     `json:"user" gorm:"foreignKey:UserID"`
	Province *Province `json:"province" gorm:"foreignKey:ProvinceCode;references:Code"`
	Regency  *Regency  `json:"regency" gorm:"foreignKey:RegencyCode;references:Code"`
	District *District `json:"district" gorm:"foreignKey:DistrictCode;references:Code"`
	Village  *Village  `json:"village" gorm:"foreignKey:VillageCode;references:Code"`
}

func (Participant) TableName() string {
	return "participants"
}
