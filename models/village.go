package models

type Village struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Code         string `json:"code" gorm:"size:20;unique;not null"`
	Name         string `json:"name" gorm:"size:100;not null"`
	DistrictCode string `json:"district_code" gorm:"size:15"`

	District *District `json:"district" gorm:"foreignKey:DistrictCode;references:Code"`
}

func (Village) TableName() string {
	return "villages"
}
