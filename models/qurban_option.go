package models

import (
	"time"

	"gorm.io/gorm"
)

type QurbanAnimalType string
type QurbanSchemeType string

const (
	Cow        QurbanAnimalType = "cow"
	Goat       QurbanAnimalType = "goat"
	Group      QurbanSchemeType = "group"
	Individual QurbanSchemeType = "individual"
)

type QurbanOption struct {
	ID             uint             `json:"id" gorm:"primaryKey"`
	QurbanPeriodID uint             `json:"qurban_period_id"`
	AnimalType     QurbanAnimalType `json:"animal_type" gorm:"type:enum('cow','goat');not null"`
	SchemeType     QurbanSchemeType `json:"scheme_type" gorm:"type:enum('group','individual');not null"`
	Price          float64          `json:"price"`
	Slots          int              `json:"slots" gorm:"default:1"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
	DeletedAt      gorm.DeletedAt   `json:"deleted_at" gorm:"index"`

	QurbanPeriod *QurbanPeriod `json:"qurban_period" gorm:"foreignKey:QurbanPeriodID"`
}

func (QurbanOption) TableName() string {
	return "qurban_options"
}
