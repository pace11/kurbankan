package model

import "time"

type Participant struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	UserID       uint      `json:"user_id"`
	Name         string    `json:"name"`
	Address      *string   `json:"address"`
	ProvinceCode string    `json:"province_code"`
	RegencyCode  string    `json:"regency_code"`
	DistrictCode string    `json:"district_code"`
	VillageCode  string    `json:"village_code"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	User     *User     `json:"user" gorm:"foreignKey:UserID"`
	Province *Province `json:"province" gorm:"foreignKey:ProvinceCode;references:Code"`
	Regency  *Regency  `json:"regency" gorm:"foreignKey:RegencyCode;references:Code"`
	District *District `json:"district" gorm:"foreignKey:DistrictCode;references:Code"`
	Village  *Village  `json:"village" gorm:"foreignKey:VillageCode;references:Code"`
}

func (Participant) TableName() string {
	return "participants"
}
