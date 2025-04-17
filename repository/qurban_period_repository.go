package repository

import (
	"kurbankan/config"
	"kurbankan/models"
)

type QurbanPeriodRepository interface {
	Index() []models.QurbanPeriod
	Save(qurbanperiod *models.QurbanPeriod)
	Update(id uint, qurbanperiod *models.QurbanPeriod) bool
	Delete(id uint) bool
}

type qurbanPeriodRepo struct{}

func NewQurbanPeriodRepository() QurbanPeriodRepository {
	return &qurbanPeriodRepo{}
}

func (r *qurbanPeriodRepo) Index() []models.QurbanPeriod {
	var qurbanPeriods []models.QurbanPeriod
	config.DB.Find(&qurbanPeriods)
	return qurbanPeriods
}

func (r *qurbanPeriodRepo) Save(qurbanperiod *models.QurbanPeriod) {
	config.DB.Create(qurbanperiod)
}

func (r *qurbanPeriodRepo) Update(id uint, qurbanperiod *models.QurbanPeriod) bool {
	var existing models.QurbanPeriod
	if err := config.DB.First(&existing, id).Error; err != nil {
		return false
	}
	existing.Year = qurbanperiod.Year
	existing.StartDate = qurbanperiod.StartDate
	existing.EndDate = qurbanperiod.EndDate
	existing.Description = qurbanperiod.Description
	config.DB.Save(&existing)
	return true
}

func (r *qurbanPeriodRepo) Delete(id uint) bool {
	result := config.DB.Delete(&models.QurbanPeriod{}, id)
	return result.RowsAffected > 0
}
