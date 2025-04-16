package model

import "time"

type QurbanDistribution struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	QurbanPeriodID uint      `json:"qurban_period_id"`
	BeneficiaryID  uint      `json:"beneficiary_id"`
	MosqueID       uint      `json:"mosque_id"`
	Amount         float64   `json:"amount"`
	Note           *string   `json:"note"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	QurbanPeriod *QurbanPeriod `json:"qurban_period" gorm:"foreignKey:QurbanPeriodID"`
	Beneficiary  *Beneficiary  `json:"beneficiary" gorm:"foreignKey:BeneficiaryID"`
	Mosque       *Mosque       `json:"mosque" gorm:"foreignKey:MosqueID"`
}

func (QurbanDistribution) TableName() string {
	return "qurban_distributions"
}
