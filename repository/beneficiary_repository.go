package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BeneficiaryRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.BeneficiaryResponse, int, any, int64, int, int)
	Show(id uint) (*models.BeneficiaryResponse, int, string, map[string]string)
	Save(beneficiary *models.Beneficiary) (any, int, string, map[string]string)
	Update(id uint, beneficiary *models.Beneficiary) (any, int, string, map[string]string)
	Delete(id uint) (any, int, string, map[string]string)
}

type beneficiaryRepo struct{}

func NewBeneficiaryRepository() BeneficiaryRepository {
	return &beneficiaryRepo{}
}

func (r *beneficiaryRepo) Index(c *gin.Context, filters map[string]any) ([]models.BeneficiaryResponse, int, any, int64, int, int) {
	var beneficiaries []models.Beneficiary
	var total int64

	query := utils.FilterByParams(config.DB.Model(&models.Beneficiary{}), filters)
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&beneficiaries)

	var response []models.BeneficiaryResponse
	for _, b := range beneficiaries {
		response = append(response, models.BeneficiaryResponse{
			ID:        b.ID,
			MosqueID:  b.MosqueID,
			Name:      b.Name,
			Address:   b.Address,
			Phone:     b.Phone,
			CreatedAt: b.CreatedAt,
			UpdatedAt: b.UpdatedAt,
		})
	}

	return response, http.StatusOK, "beneficiary", total, page, limit
}

func (r *beneficiaryRepo) Show(id uint) (*models.BeneficiaryResponse, int, string, map[string]string) {
	var beneficiary models.Beneficiary

	if err := config.DB.Where("id = ?", id).First(&beneficiary).Error; err != nil {
		return nil, http.StatusNotFound, "beneficiary", nil
	}

	response := &models.BeneficiaryResponse{
		ID:        beneficiary.ID,
		MosqueID:  beneficiary.MosqueID,
		Name:      beneficiary.Name,
		Address:   beneficiary.Address,
		Phone:     beneficiary.Phone,
		CreatedAt: beneficiary.CreatedAt,
		UpdatedAt: beneficiary.UpdatedAt,
	}

	return response, http.StatusOK, "beneficiary", nil
}

func (r *beneficiaryRepo) Save(beneficiary *models.Beneficiary) (any, int, string, map[string]string) {
	if err := config.DB.Create(beneficiary).Error; err != nil {
		return nil, http.StatusInternalServerError, "beneficiary", nil
	}

	return beneficiary, http.StatusCreated, "beneficiary", nil
}

func (r *beneficiaryRepo) Update(id uint, beneficiary *models.Beneficiary) (any, int, string, map[string]string) {
	var existing models.Beneficiary
	if err := config.DB.First(&existing, id).Error; err != nil {
		return nil, http.StatusNotFound, "beneficiary", nil
	}

	existing.MosqueID = beneficiary.MosqueID
	existing.Name = beneficiary.Name
	existing.Address = beneficiary.Address
	existing.Phone = beneficiary.Phone

	if err := config.DB.Save(&existing).Error; err != nil {
		return nil, http.StatusInternalServerError, "beneficiary", nil
	}

	return beneficiary, http.StatusOK, "beneficiary", nil
}

func (r *beneficiaryRepo) Delete(id uint) (any, int, string, map[string]string) {
	result := config.DB.Delete(&models.Beneficiary{}, id)

	if result.Error != nil {
		return nil, http.StatusInternalServerError, "beneficiary", nil
	}

	if result.RowsAffected == 0 {
		return nil, http.StatusNotFound, "beneficiary", nil
	}

	return nil, http.StatusOK, "beneficiary", nil
}
