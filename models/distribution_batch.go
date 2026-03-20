package models

import "time"

type DistributionBatch struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	MosqueID       uint      `json:"mosque_id"`
	QurbanPeriodID uint      `json:"qurban_period_id"`
	TotalPackages  int       `json:"total_packages"`
	CreatedAt      time.Time `json:"created_at"`

	Mosque       *Mosque       `json:"mosque" gorm:"foreignKey:MosqueID"`
	QurbanPeriod *QurbanPeriod `json:"qurban_period" gorm:"foreignKey:QurbanPeriodID"`
}

func (DistributionBatch) TableName() string {
	return "distribution_batches"
}
