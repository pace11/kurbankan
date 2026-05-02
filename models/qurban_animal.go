package models

import (
	"encoding/json"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AnimalStatus string

const (
	AnimalPurchased   AnimalStatus = "purchased"
	AnimalArrived     AnimalStatus = "arrived"
	AnimalSlaughtered AnimalStatus = "slaughtered"
	AnimalProcessed   AnimalStatus = "processed"
	AnimalDistributed AnimalStatus = "distributed"
)

type StatusLogEntry struct {
	Status    AnimalStatus `json:"status"`
	Timestamp time.Time    `json:"timestamp"`
}

type QurbanAnimal struct {
	ID uint `json:"id" gorm:"primaryKey"`

	MosqueID       uint `json:"mosque_id"`
	QurbanPeriodID uint `json:"qurban_period_id"`

	Status        AnimalStatus   `json:"status" gorm:"type:varchar(20);default:'purchased'"`
	StatusHistory datatypes.JSON `json:"status_history" gorm:"type:jsonb"` // Auto-updated
	SlaughteredAt *time.Time     `json:"slaughtered_at"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Mosque       *Mosque       `json:"mosque" gorm:"foreignKey:MosqueID"`
	QurbanPeriod *QurbanPeriod `json:"qurban_period" gorm:"foreignKey:QurbanPeriodID"`
}

func (QurbanAnimal) TableName() string {
	return "qurban_animals"
}

// Hook: auto-update StatusHistory when status changes
func (q *QurbanAnimal) BeforeUpdate(tx *gorm.DB) error {
	// Check apakah field Status berubah
	if tx.Statement.Changed("Status") {
		// Parse existing history
		var history []StatusLogEntry
		if q.StatusHistory != nil {
			json.Unmarshal(q.StatusHistory, &history)
		}

		// Append new log entry
		history = append(history, StatusLogEntry{
			Status:    q.Status,
			Timestamp: time.Now(),
		})

		// Save back to JSONB field
		historyJSON, _ := json.Marshal(history)
		q.StatusHistory = historyJSON
	}
	return nil
}

// Hook untuk create (log pertama kali)
func (q *QurbanAnimal) AfterCreate(tx *gorm.DB) error {
	history := []StatusLogEntry{
		{
			Status:    q.Status,
			Timestamp: time.Now(),
		},
	}

	historyJSON, _ := json.Marshal(history)
	q.StatusHistory = historyJSON

	// Update field StatusHistory
	tx.Model(q).Update("status_history", historyJSON)
	return nil
}
