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
	Save(transaction *models.TransactionPayload) (any, int, string, map[string]string)
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
			ID:             t.ID,
			IsFull:         t.IsFull,
			CreatedAt:      t.CreatedAt,
			UpdatedAt:      t.UpdatedAt,
			Mosque:         t.Mosque,
			QurbanPeriod:   t.QurbanPeriod,
			QurbanOffering: t.QurbanOffering,
		})
	}

	return response, http.StatusOK, "transaction", total, page, limit
}

func (r *transactionRepo) Save(transaction *models.TransactionPayload) (any, int, string, map[string]string) {
	var qurbanOffering models.QurbanOffering

	if err := config.DB.First(&qurbanOffering, transaction.QurbanOfferingID).Error; err != nil {
		return nil, http.StatusNotFound, "qurban offering", nil
	}

	transactionData := models.Transaction{
		QurbanPeriodID:   qurbanOffering.QurbanPeriodID,
		MosqueID:         transaction.MosqueID,
		QurbanOfferingID: transaction.QurbanOfferingID,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if qurbanOffering.AnimalType == models.Cow && qurbanOffering.SchemeType == models.Group {
		var itemCount int64
		config.DB.Model(&models.TransactionItem{}).
			Joins("JOIN transactions ON transaction_items.transaction_id = transactions.id").
			Where("transactions.qurban_offering_id = ? AND transaction_items.status != ?", qurbanOffering.ID, models.Cancelled).
			Count(&itemCount)
	} else {
		transactionData.IsFull = true
	}

	return transactionData, http.StatusCreated, "transaction", nil
}
