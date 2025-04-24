package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DistrictRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.District, int, any, any, int64, int, int)
}

type districtRepository struct{}

func NewDistrictRepository() DistrictRepository {
	return &districtRepository{}
}

func (r *districtRepository) Index(c *gin.Context, filters map[string]any) ([]models.District, int, any, any, int64, int, int) {
	var districts []models.District
	var total int64

	query := utils.FilterByParams(config.DB.Model(&models.District{}), filters)
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&districts)
	return districts, http.StatusOK, "district", "get", total, page, limit
}
