package controllers

import (
	"fmt"
	"net/http"
	"time"

	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateTransactionRequest struct {
	MosqueID       uint    `json:"mosque_id" binding:"required"`
	QurbanOptionID uint    `json:"qurban_option_id" binding:"required"`
	ParticipantID  uint    `json:"participant_id" binding:"required"`
	Amount         float64 `json:"amount" binding:"required"`
}

func CreateTransaction(c *gin.Context) {
	var payload CreateTransactionRequest
	if utils.BindAndValidate(c, &payload) != nil {
		return
	}

	var option models.QurbanOption
	if err := config.DB.First(&option, payload.QurbanOptionID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Qurban option not found")
		return
	}

	if option.AnimalType == models.Cow && option.SchemeType == models.Group {
		var itemCount int64
		config.DB.Model(&models.TransactionItem{}).
			Joins("JOIN transactions ON transaction_items.transaction_id = transactions.id").
			Where("transactions.qurban_option_id = ? AND transaction_items.status != ?", option.ID, models.Cancelled).
			Count(&itemCount)

		if itemCount >= int64(option.Slots) {
			utils.ErrorResponse(c, http.StatusBadRequest, "Quota for this group cow qurban is full")
			return
		}
	}

	tx := config.DB.Begin()
	transaction := models.Transaction{
		QurbanPeriodID: option.QurbanPeriodID,
		MosqueID:       payload.MosqueID,
		QurbanOptionID: payload.QurbanOptionID,
		IsFull:         false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create transaction")
		return
	}

	externalID := fmt.Sprintf("qurban-%s", uuid.NewString())

	item := models.TransactionItem{
		TransactionID: transaction.ID,
		ParticipantID: payload.ParticipantID,
		Amount:        payload.Amount,
		Status:        models.Pending,
		PaymentType:   models.VA,
		ExternalID:    externalID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := tx.Create(&item).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create transaction item")
		return
	}

	va, err := utils.CreateVirtualAccount(externalID, "BNI", "Qurban User", payload.Amount)
	if err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create virtual account")
		return
	}

	tx.Commit()
	c.JSON(http.StatusCreated, gin.H{
		"transaction": transaction,
		"item":        item,
		"va":          va,
	})
}
