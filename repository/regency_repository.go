package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegencyRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.Regency, int, any, any, int64, int, int)
}

type regencyRepository struct{}

func NewRegencyRepository() RegencyRepository {
	return &regencyRepository{}
}

func (r *regencyRepository) Index(c *gin.Context, filters map[string]any) ([]models.Regency, int, any, any, int64, int, int) {
	var regencies []models.Regency
	var total int64

	query := utils.FilterByParams(config.DB.Model(&models.Regency{}), filters)
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&regencies)
	return regencies, http.StatusOK, "regency", "get", total, page, limit
}
