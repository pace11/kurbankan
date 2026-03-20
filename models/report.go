package models

import (
"time"

"gorm.io/datatypes"
)

type Report struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	MosqueID       uint           `json:"mosque_id"`
	QurbanPeriodID uint           `json:"qurban_period_id"`
	GeneratedBy    uint           `json:"generated_by"`
	SnapshotJSON   datatypes.JSON `json:"snapshot_json" gorm:"type:jsonb"`
	CreatedAt      time.Time      `json:"created_at"`

	Mosque        *Mosque       `json:"mosque" gorm:"foreignKey:MosqueID"`
	QurbanPeriod  *QurbanPeriod `json:"qurban_period" gorm:"foreignKey:QurbanPeriodID"`
	GeneratedUser *User         `json:"generated_user" gorm:"foreignKey:GeneratedBy"`
}

func (Report) TableName() string {
	return "reports"
}
