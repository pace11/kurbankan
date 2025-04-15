package model

type Province struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"size:10;unique;not null"`
	Name string `json:"name" gorm:"size:100;not null"`
}

func (Province) TableName() string {
	return "provinces"
}
