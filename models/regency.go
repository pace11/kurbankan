package models

type Regency struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Code         string `json:"code" gorm:"size:10;unique;not null"`
	Name         string `json:"name" gorm:"size:100;not null"`
	ProvinceCode string `json:"province_code" gorm:"size:10"`

	Province *Province `json:"province" gorm:"foreignKey:ProvinceCode;references:Code"`
}

func (Regency) TableName() string {
	return "regencies"
}
