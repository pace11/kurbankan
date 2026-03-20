package models

import "time"

type MediaType string

const (
MediaPhoto MediaType = "photo"
MediaVideo MediaType = "video"
)

type AnimalMedia struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	AnimalID  uint      `json:"animal_id"`
	URL       string    `json:"url" gorm:"type:text"`
	Type      MediaType `json:"type" gorm:"type:varchar(20)"`
	CreatedAt time.Time `json:"created_at"`

	Animal *QurbanAnimal `json:"animal" gorm:"foreignKey:AnimalID"`
}

func (AnimalMedia) TableName() string {
	return "animal_media"
}
