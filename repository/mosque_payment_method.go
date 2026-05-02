package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"net/http"
)

type MosquePaymentMethodRepository interface {
	// Index(c *gin.Context, mosqueID uint) ([]models.MosquePaymentMethod, int, any, int64, int, int)
	Create(payload *models.MosquePaymentMethodCreatePayload) (any, int, string, map[string]string)
}

type mosquePaymentMethodRepo struct{}

func NewMosquePaymentMethodRepository() MosquePaymentMethodRepository {
	return &mosquePaymentMethodRepo{}
}

func (r *mosquePaymentMethodRepo) Create(payload *models.MosquePaymentMethodCreatePayload) (any, int, string, map[string]string) {
	mosquePaymentMethod := models.MosquePaymentMethod{
		MosqueID:      payload.MosqueID,
		BankID:        payload.BankID,
		Type:          payload.Type,
		Name:          payload.Name,
		AccountName:   payload.AccountName,
		AccountNumber: payload.AccountNumber,
		QRISImageURL:  payload.QRISImageURL,
		Instructions:  payload.Instructions,
	}

	if err := config.DB.Create(&mosquePaymentMethod).Error; err != nil {
		return nil, http.StatusInternalServerError, "mosque payment method", nil
	}

	return mosquePaymentMethod, http.StatusCreated, "mosque payment method", nil
}
