package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QurbanOptionRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.QurbanOptionResponse, int, any, int64, int, int)
	Save(qurbanoption *models.QurbanOption) (any, int, string, map[string]string)
	Update(id uint, qurbanoption *models.QurbanOption) (any, int, string, map[string]string)
	Delete(id uint) (any, int, string, map[string]string)
}

type qurbanOptionRepo struct{}

func NewQurbanOptionRepository() QurbanOptionRepository {
	return &qurbanOptionRepo{}
}

func (r *qurbanOptionRepo) Index(c *gin.Context, filters map[string]any) ([]models.QurbanOptionResponse, int, any, int64, int, int) {
	var qurbanOptions []models.QurbanOption
	var total int64

	query := utils.FilterByParams(config.DB.Model(&models.QurbanOption{}), filters)
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&qurbanOptions)

	var response []models.QurbanOptionResponse
	for _, q := range qurbanOptions {
		response = append(response, models.QurbanOptionResponse{
			ID:             q.ID,
			QurbanPeriodID: q.QurbanPeriodID,
			AnimalType:     q.AnimalType,
			SchemeType:     q.SchemeType,
			Price:          q.Price,
			Slots:          q.Slots,
			CreatedAt:      q.CreatedAt,
			UpdatedAt:      q.UpdatedAt,
		})
	}

	return response, http.StatusOK, "qurban option", total, page, limit
}

func (r *qurbanOptionRepo) Save(qurbanoption *models.QurbanOption) (any, int, string, map[string]string) {
	if err := config.DB.Create(qurbanoption).Error; err != nil {
		return nil, http.StatusInternalServerError, "qurban option", nil
	}

	return qurbanoption, http.StatusCreated, "qurban option", nil
}

func (r *qurbanOptionRepo) Update(id uint, qurbanoption *models.QurbanOption) (any, int, string, map[string]string) {
	var existing models.QurbanOption

	if err := config.DB.First(&existing, id).Error; err != nil {
		return nil, http.StatusNotFound, "qurban option", nil
	}

	if err := config.DB.Model(&existing).Updates(map[string]any{
		"qurban_period_id": qurbanoption.QurbanPeriodID,
		"animal_type":      qurbanoption.AnimalType,
		"scheme_type":      qurbanoption.SchemeType,
		"price":            qurbanoption.Price,
		"slots":            qurbanoption.Slots,
	}).Error; err != nil {
		return nil, http.StatusInternalServerError, "qurban option", nil
	}

	return qurbanoption, http.StatusOK, "qurban option", nil
}

func (r *qurbanOptionRepo) Delete(id uint) (any, int, string, map[string]string) {
	result := config.DB.Delete(&models.QurbanOption{}, id)

	if result.Error != nil {
		return nil, http.StatusInternalServerError, "qurban option", nil
	}

	if result.RowsAffected == 0 {
		return nil, http.StatusNotFound, "qurban option", nil
	}

	return nil, http.StatusOK, "qurban option", nil
}
