package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QurbanOfferingRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.QurbanOfferingResponse, int, any, int64, int, int)
	Save(qurbanOffering *models.QurbanOffering) (any, int, string, map[string]string)
	Update(id uint, qurbanOffering *models.QurbanOffering) (any, int, string, map[string]string)
	Delete(id uint) (any, int, string, map[string]string)
}

type qurbanOfferingRepo struct{}

func NewQurbanOfferingRepository() QurbanOfferingRepository {
	return &qurbanOfferingRepo{}
}

func (r *qurbanOfferingRepo) Index(c *gin.Context, filters map[string]any) ([]models.QurbanOfferingResponse, int, any, int64, int, int) {
	var qurbanOfferings []models.QurbanOffering
	var total int64

	query := utils.FilterByParams(config.DB.Model(&models.QurbanOffering{}), filters)
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&qurbanOfferings)

	var response []models.QurbanOfferingResponse
	for _, q := range qurbanOfferings {
		response = append(response, models.QurbanOfferingResponse{
			ID:             q.ID,
			QurbanPeriodID: q.QurbanPeriodID,
			AnimalType:     q.AnimalType,
			SchemeType:     q.SchemeType,
			Price:          q.Price,
			Capacity:       q.Capacity,
			FilledSlots:    q.FilledSlots,
			CreatedAt:      q.CreatedAt,
			UpdatedAt:      q.UpdatedAt,
		})
	}

	return response, http.StatusOK, "qurban offering", total, page, limit
}

func (r *qurbanOfferingRepo) Save(qurbanOffering *models.QurbanOffering) (any, int, string, map[string]string) {
	if err := config.DB.Create(qurbanOffering).Error; err != nil {
		return nil, http.StatusInternalServerError, "qurban offering", nil
	}

	return qurbanOffering, http.StatusCreated, "qurban offering", nil
}

func (r *qurbanOfferingRepo) Update(id uint, qurbanOffering *models.QurbanOffering) (any, int, string, map[string]string) {
	var existing models.QurbanOffering

	if err := config.DB.First(&existing, id).Error; err != nil {
		return nil, http.StatusNotFound, "qurban offering", nil
	}

	if err := config.DB.Model(&existing).Updates(map[string]any{
		"qurban_period_id": qurbanOffering.QurbanPeriodID,
		"animal_type":      qurbanOffering.AnimalType,
		"scheme_type":      qurbanOffering.SchemeType,
		"price":            qurbanOffering.Price,
		"capacity":         qurbanOffering.Capacity,
		"filled_slots":     qurbanOffering.FilledSlots,
	}).Error; err != nil {
		return nil, http.StatusInternalServerError, "qurban offering", nil
	}

	return qurbanOffering, http.StatusOK, "qurban offering", nil
}

func (r *qurbanOfferingRepo) Delete(id uint) (any, int, string, map[string]string) {
	result := config.DB.Delete(&models.QurbanOffering{}, id)

	if result.Error != nil {
		return nil, http.StatusInternalServerError, "qurban option", nil
	}

	if result.RowsAffected == 0 {
		return nil, http.StatusNotFound, "qurban option", nil
	}

	return nil, http.StatusOK, "qurban option", nil
}
