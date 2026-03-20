package models

import "time"

type AnimalParticipant struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	AnimalID          uint      `json:"animal_id"`
	ParticipantID     uint      `json:"participant_id"`
	TransactionItemID uint      `json:"transaction_item_id"`
	CreatedAt         time.Time `json:"created_at"`

	Animal          *QurbanAnimal    `json:"animal" gorm:"foreignKey:AnimalID"`
	Participant     *Participant     `json:"participant" gorm:"foreignKey:ParticipantID"`
	TransactionItem *TransactionItem `json:"transaction_item" gorm:"foreignKey:TransactionItemID"`
}

func (AnimalParticipant) TableName() string {
	return "animal_participants"
}
