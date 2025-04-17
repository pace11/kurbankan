package models

type District struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Code        string `json:"code" gorm:"size:15;unique;not null"`
	Name        string `json:"name" gorm:"size:100;not null"`
	RegencyCode string `json:"regency_code" gorm:"size:10"`

	Regency *Regency `json:"regency" gorm:"foreignKey:RegencyCode;references:Code"`
}

func (District) TableName() string {
	return "districts"
}
