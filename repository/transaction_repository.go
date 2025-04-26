package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionRepository interface {
	Index(c *gin.Context) ([]models.TransactionResponse, int, any, int64, int, int)
	Save(transaction *models.TransactionDTO) (any, int, string, map[string]string)
}

type transactionRepo struct{}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepo{}
}

func (r *transactionRepo) Index(c *gin.Context) ([]models.TransactionResponse, int, any, int64, int, int) {
	var transactions []models.Transaction
	var total int64

	query := config.DB.Model(&models.Transaction{}).Preload("Mosque").Preload("QurbanPeriod").Preload("QurbanOption")
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&transactions)

	var response []models.TransactionResponse
	for _, t := range transactions {
		response = append(response, models.TransactionResponse{
			ID:           t.ID,
			IsFull:       t.IsFull,
			CreatedAt:    t.CreatedAt,
			UpdatedAt:    t.UpdatedAt,
			Mosque:       t.Mosque,
			QurbanPeriod: t.QurbanPeriod,
			QurbanOption: t.QurbanOption,
		})
	}

	return response, http.StatusOK, "transaction", total, page, limit
}

func (r *transactionRepo) Save(transaction *models.TransactionDTO) (any, int, string, map[string]string) {
	var qurbanOption models.QurbanOption

	if err := config.DB.First(&qurbanOption, transaction.QurbanOptionID).Error; err != nil {
		return nil, http.StatusNotFound, "qurban option", nil
	}

	transactionData := models.Transaction{
		QurbanPeriodID: qurbanOption.QurbanPeriodID,
		MosqueID:       transaction.MosqueID,
		QurbanOptionID: transaction.QurbanOptionID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if qurbanOption.AnimalType == models.Cow && qurbanOption.SchemeType == models.Group {
		var itemCount int64
		config.DB.Model(&models.TransactionItem{}).
			Joins("JOIN transactions ON transaction_items.transaction_id = transactions.id").
			Where("transactions.qurban_option_id = ? AND transaction_items.status != ?", qurbanOption.ID, models.Cancelled).
			Count(&itemCount)
	} else {
		transactionData.IsFull = true
	}

	return transactionData, http.StatusCreated, "transaction", nil
}
