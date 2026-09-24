package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QurbanOfferingRepository interface {
	ListWithPagination(c *gin.Context, mosqueID uint, filters map[string]any) ([]models.QurbanOfferingResponse, int64, int, int)
	Save(payload *models.QurbanOfferingRequest) (any, int, string, map[string]string)
	Update(payload *models.QurbanOfferingRequest) (any, int, string, map[string]string)
	Delete(id uint) (any, int, string, map[string]string)
}

type qurbanOfferingRepo struct{}

func NewQurbanOfferingRepository() QurbanOfferingRepository {
	return &qurbanOfferingRepo{}
}

func (r *qurbanOfferingRepo) ListWithPagination(c *gin.Context, mosqueID uint, filters map[string]any) ([]models.QurbanOfferingResponse, int64, int, int) {
	var qurbanOfferings []models.QurbanOffering
	var total int64

	query := utils.FilterByParams(config.DB.Model(&models.QurbanOffering{}).Preload("Mosque").Where("mosque_id = ?", mosqueID), filters)
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
			Name:           q.Name,
			Price:          q.Price,
			Capacity:       q.Capacity,
			FilledSlots:    q.FilledSlots,
			ConfirmedSlots: q.ConfirmedSlots,
			Status:         q.Status,
			CreatedAt:      q.CreatedAt,
			UpdatedAt:      q.UpdatedAt,
		})
	}

	return response, total, page, limit
}

func (r *qurbanOfferingRepo) Save(payload *models.QurbanOfferingRequest) (any, int, string, map[string]string) {
	qurbanOffering := &models.QurbanOffering{
		QurbanPeriodID: payload.QurbanPeriodID,
		AnimalType:     payload.AnimalType,
		SchemeType:     payload.SchemeType,
		Name:           payload.Name,
		Price:          payload.Price,
		Capacity:       payload.Capacity,
		Status:         payload.Status,
		MosqueID:       payload.MosqueID,
	}

	if err := config.DB.Create(qurbanOffering).Error; err != nil {
		return nil, http.StatusInternalServerError, "qurban offering", nil
	}

	return qurbanOffering, http.StatusCreated, "qurban offering", nil
}

func (r *qurbanOfferingRepo) Update(payload *models.QurbanOfferingRequest) (any, int, string, map[string]string) {
	var existing models.QurbanOffering

	if err := config.DB.First(&existing, payload.ID).Error; err != nil {
		return nil, http.StatusNotFound, "qurban offering", nil
	}

	// Update fields
	existing.QurbanPeriodID = payload.QurbanPeriodID
	existing.AnimalType = payload.AnimalType
	existing.SchemeType = payload.SchemeType
	existing.Name = payload.Name
	existing.Price = payload.Price
	existing.Capacity = payload.Capacity
	existing.Status = payload.Status

	if err := config.DB.Model(&existing).Updates(&existing).Error; err != nil {
		return nil, http.StatusInternalServerError, "qurban offering", nil
	}

	return &existing, http.StatusOK, "qurban offering", nil
}

func (r *qurbanOfferingRepo) Delete(id uint) (any, int, string, map[string]string) {
	result := config.DB.Delete(&models.QurbanOffering{}, id)

	if result.Error != nil {
		return nil, http.StatusInternalServerError, "qurban offering", nil
	}

	if result.RowsAffected == 0 {
		return nil, http.StatusNotFound, "qurban offering", nil
	}

	return nil, http.StatusOK, "qurban offering", nil
}
