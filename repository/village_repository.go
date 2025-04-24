package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VillageRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.Village, int, any, int64, int, int)
}

type villageRepository struct{}

func NewVillageRepository() VillageRepository {
	return &villageRepository{}
}

func (r *villageRepository) Index(c *gin.Context, filters map[string]any) ([]models.Village, int, any, int64, int, int) {
	var villages []models.Village
	var total int64

	query := utils.FilterByParams(config.DB.Model(&models.Village{}), filters)
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&villages)
	return villages, http.StatusOK, "village", total, page, limit
}
