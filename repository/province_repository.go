package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"

	"github.com/gin-gonic/gin"
)

type ProvinceRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.Province, int64, int, int)
}

type provinceRepository struct{}

func NewProvinceRepository() ProvinceRepository {
	return &provinceRepository{}
}

func (r *provinceRepository) Index(c *gin.Context, filters map[string]any) ([]models.Province, int64, int, int) {
	var provinces []models.Province
	var total int64

	query := utils.FilterByParams(config.DB.Model(&models.Province{}), filters)
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&provinces)
	return provinces, total, page, limit
}
