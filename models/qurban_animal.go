package models

import "time"

type AnimalStatus string

const (
	AnimalPurchased   AnimalStatus = "purchased"
	AnimalArrived     AnimalStatus = "arrived"
	AnimalSlaughtered AnimalStatus = "slaughtered"
	AnimalProcessed   AnimalStatus = "processed"
	AnimalDistributed AnimalStatus = "distributed"
)

type QurbanAnimal struct {
	ID             uint             `json:"id" gorm:"primaryKey"`
	MosqueID       uint             `json:"mosque_id"`
	QurbanPeriodID uint             `json:"qurban_period_id"`
	Type           QurbanAnimalType `json:"type" gorm:"type:varchar(20)"`
	Name           string           `json:"name" gorm:"size:100"`
	Weight         float64          `json:"weight" gorm:"type:decimal(6,2)"`
	Price          float64          `json:"price" gorm:"type:decimal(12,2)"`
	Status         AnimalStatus     `json:"status" gorm:"type:varchar(20);default:'purchased'"`
	SlaughteredAt  *time.Time       `json:"slaughtered_at"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`

	Mosque       *Mosque       `json:"mosque" gorm:"foreignKey:MosqueID"`
	QurbanPeriod *QurbanPeriod `json:"qurban_period" gorm:"foreignKey:QurbanPeriodID"`
}

func (QurbanAnimal) TableName() string {
	return "qurban_animals"
}
