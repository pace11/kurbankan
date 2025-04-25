package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QurbanPeriodRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.QurbanPeriodResponse, int, any, int64, int, int)
	Save(qurbanperiod *models.QurbanPeriod) (any, int, string, map[string]string)
	Update(id uint, qurbanperiod *models.QurbanPeriod) (any, int, string, map[string]string)
	Delete(id uint) (any, int, string, map[string]string)
}

type qurbanPeriodRepo struct{}

func NewQurbanPeriodRepository() QurbanPeriodRepository {
	return &qurbanPeriodRepo{}
}

func (r *qurbanPeriodRepo) Index(c *gin.Context, filters map[string]any) ([]models.QurbanPeriodResponse, int, any, int64, int, int) {
	var qurbanPeriods []models.QurbanPeriod
	var total int64

	query := utils.FilterByParams(config.DB.Model(&models.QurbanPeriod{}), filters)
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&qurbanPeriods)

	var response []models.QurbanPeriodResponse
	for _, q := range qurbanPeriods {
		response = append(response, models.QurbanPeriodResponse{
			ID:          q.ID,
			Year:        q.Year,
			StartDate:   q.StartDate,
			EndDate:     q.EndDate,
			Description: q.Description,
			CreatedAt:   q.CreatedAt,
			UpdatedAt:   q.UpdatedAt,
		})
	}

	return response, http.StatusOK, "qurban period", total, page, limit
}

func (r *qurbanPeriodRepo) Save(qurbanperiod *models.QurbanPeriod) (any, int, string, map[string]string) {
	if err := config.DB.Create(qurbanperiod).Error; err != nil {
		return nil, http.StatusInternalServerError, "qurban period", nil
	}

	return qurbanperiod, http.StatusCreated, "qurban period", nil
}

func (r *qurbanPeriodRepo) Update(id uint, qurbanperiod *models.QurbanPeriod) (any, int, string, map[string]string) {
	var existing models.QurbanPeriod

	if err := config.DB.First(&existing, id).Error; err != nil {
		return nil, http.StatusNotFound, "qurban period", nil
	}

	if err := config.DB.Model(&existing).Updates(map[string]any{
		"year":        qurbanperiod.Year,
		"start_date":  qurbanperiod.StartDate,
		"end_date":    qurbanperiod.EndDate,
		"description": qurbanperiod.Description,
	}).Error; err != nil {
		return nil, http.StatusInternalServerError, "qurban period", nil
	}

	return qurbanperiod, http.StatusOK, "qurban period", nil
}

func (r *qurbanPeriodRepo) Delete(id uint) (any, int, string, map[string]string) {
	result := config.DB.Delete(&models.QurbanPeriod{}, id)

	if result.Error != nil {
		return nil, http.StatusInternalServerError, "qurban period", nil
	}

	if result.RowsAffected == 0 {
		return nil, http.StatusNotFound, "qurban period", nil
	}

	return nil, http.StatusOK, "qurban period", nil
}
