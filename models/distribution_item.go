package models

import "time"

type DistributionItem struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	BatchID       uint       `json:"batch_id"`
	BeneficiaryID uint       `json:"beneficiary_id"`
	ReceivedAt    *time.Time `json:"received_at"`
	Notes         *string    `json:"notes" gorm:"type:text"`
	CreatedAt     time.Time  `json:"created_at"`

	Batch       *DistributionBatch `json:"batch" gorm:"foreignKey:BatchID"`
	Beneficiary *Beneficiary       `json:"beneficiary" gorm:"foreignKey:BeneficiaryID"`
}

func (DistributionItem) TableName() string {
	return "distribution_items"
}
