package models

import "time"

type AnimalLog struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	AnimalID  uint      `json:"animal_id"`
	Status    string    `json:"status" gorm:"size:50"`
	Note      *string   `json:"note" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`

	Animal *QurbanAnimal `json:"animal" gorm:"foreignKey:AnimalID"`
}

func (AnimalLog) TableName() string {
	return "animal_logs"
}
