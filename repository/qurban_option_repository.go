package repository

import (
	"kurbankan/config"
	"kurbankan/models"
)

type QurbanOptionRepository interface {
	Index() []models.QurbanOption
	Save(qurbanoption *models.QurbanOption)
	Update(id uint, qurbanoption *models.QurbanOption) bool
	Delete(id uint) bool
}

type qurbanOptionRepo struct{}

func NewQurbanOptionRepository() QurbanOptionRepository {
	return &qurbanOptionRepo{}
}

func (r *qurbanOptionRepo) Index() []models.QurbanOption {
	var qurbanOptions []models.QurbanOption
	config.DB.Find(&qurbanOptions)
	return qurbanOptions
}

func (r *qurbanOptionRepo) Save(qurbanoption *models.QurbanOption) {
	config.DB.Create(qurbanoption)
}

func (r *qurbanOptionRepo) Update(id uint, qurbanoption *models.QurbanOption) bool {
	var existing models.QurbanOption
	if err := config.DB.First(&existing, id).Error; err != nil {
		return false
	}
	existing.QurbanPeriodID = qurbanoption.QurbanPeriodID
	existing.AnimalType = qurbanoption.AnimalType
	existing.SchemeType = qurbanoption.SchemeType
	existing.Price = qurbanoption.Price
	existing.Slots = qurbanoption.Slots
	config.DB.Save(&existing)
	return true
}

func (r *qurbanOptionRepo) Delete(id uint) bool {
	result := config.DB.Delete(&models.QurbanOption{}, id)
	return result.RowsAffected > 0
}
