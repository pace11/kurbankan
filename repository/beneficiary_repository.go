package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"

	"github.com/gin-gonic/gin"
)

type BeneficiaryRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.BeneficiaryResponse, int64, int, int)
	Save(beneficiary *models.Beneficiary)
	Update(id uint, beneficiary *models.Beneficiary) bool
	Delete(id uint) bool
}

type beneficiaryRepo struct{}

func NewBeneficiaryRepository() BeneficiaryRepository {
	return &beneficiaryRepo{}
}

func (r *beneficiaryRepo) Index(c *gin.Context, filters map[string]any) ([]models.BeneficiaryResponse, int64, int, int) {
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

	return response, total, page, limit
}

func (r *beneficiaryRepo) Save(beneficiary *models.Beneficiary) {
	config.DB.Create(beneficiary)
}

func (r *beneficiaryRepo) Update(id uint, beneficiary *models.Beneficiary) bool {
	var existing models.Beneficiary
	if err := config.DB.First(&existing, id).Error; err != nil {
		return false
	}

	existing.MosqueID = beneficiary.MosqueID
	existing.Name = beneficiary.Name
	existing.Address = beneficiary.Address
	existing.Phone = beneficiary.Phone

	config.DB.Save(&existing)
	return true
}

func (r *beneficiaryRepo) Delete(id uint) bool {
	result := config.DB.Delete(&models.Beneficiary{}, id)
	return result.RowsAffected > 0
}
