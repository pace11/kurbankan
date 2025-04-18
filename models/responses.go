package models

import "time"

type MosqueResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:text;not null"`
	Address   *string   `json:"address"`
	Photos    *string   `json:"photos"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Province *Province `json:"province" gorm:"foreignKey:ProvinceCode;references:Code"`
	Regency  *Regency  `json:"regency" gorm:"foreignKey:RegencyCode;references:Code"`
	District *District `json:"district" gorm:"foreignKey:DistrictCode;references:Code"`
	Village  *Village  `json:"village" gorm:"foreignKey:VillageCode;references:Code"`
}
