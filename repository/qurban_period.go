package repository

import (
	"fmt"
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QurbanPeriodRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.QurbanPeriodResponse, int, any, int64, int, int)
	Save(payload *models.QurbanPeriodRequest) (any, int, string, map[string]string)
	Update(payload *models.QurbanPeriodRequest) (any, int, string, map[string]string)
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

func (r *qurbanPeriodRepo) Save(payload *models.QurbanPeriodRequest) (any, int, string, map[string]string) {
	qurbanperiod := &models.QurbanPeriod{
		Year:        payload.Year,
		StartDate:   payload.StartDate,
		EndDate:     payload.EndDate,
		Description: payload.Description,
		MosqueID:    payload.MosqueID,
	}

	if err := config.DB.Create(qurbanperiod).Error; err != nil {
		fmt.Println("Error creating qurban period:", err)
		return nil, http.StatusInternalServerError, "qurban period", nil
	}

	return qurbanperiod, http.StatusCreated, "qurban period", nil
}

func (r *qurbanPeriodRepo) Update(payload *models.QurbanPeriodRequest) (any, int, string, map[string]string) {
	var existing models.QurbanPeriod

	if err := config.DB.First(&existing, payload.ID).Error; err != nil {
		return nil, http.StatusNotFound, "qurban period", nil
	}

	if err := config.DB.Model(&existing).Updates(map[string]any{
		"year":        payload.Year,
		"start_date":  payload.StartDate,
		"end_date":    payload.EndDate,
		"description": payload.Description,
	}).Error; err != nil {
		return nil, http.StatusInternalServerError, "qurban period", nil
	}

	return &existing, http.StatusOK, "qurban period", nil
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
