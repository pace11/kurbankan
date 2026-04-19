package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BankRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.Bank, int, any, int64, int, int)
}

type bankRepository struct{}

func NewBankRepository() BankRepository {
	return &bankRepository{}
}

func (r *bankRepository) Index(c *gin.Context, filters map[string]any) ([]models.Bank, int, any, int64, int, int) {
	var banks []models.Bank
	var total int64

	query := utils.FilterByParams(config.DB.Model(&models.Bank{}), filters)
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&banks)
	return banks, http.StatusOK, "bank", total, page, limit
}
