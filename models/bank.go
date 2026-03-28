package models

type Bank struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Code      string `json:"code" gorm:"size:10;unique;not null"`
	Name      string `json:"name" gorm:"size:100;not null"`
	Year      int    `json:"year" gorm:"not null"`
	Type      string `json:"type" gorm:"size:20;not null"`
	Category  string `json:"category" gorm:"size:50;not null"`
	Status    string `json:"status" gorm:"size:20;not null"`
	Ownership string `json:"ownership" gorm:"size:200"`
}

func (Bank) TableName() string {
	return "banks"
}
