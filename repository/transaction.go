package repository

import (
	"fmt"
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionRepository interface {
	Index(c *gin.Context) ([]models.TransactionResponse, int, any, int64, int, int)
	Save(transaction *models.TransactionCreatePayload) (any, int, string, map[string]string)
	UpdateProof(payload *models.TransactionUploadProofPayload) (any, int, string, map[string]string)

	// Specific by mosque
	IndexByMosqueID(c *gin.Context, mosqueID uint) ([]models.TransactionResponse, int, any, int64, int, int)

	// Specific by participant
	IndexByParticipantID(c *gin.Context, participantID uint) ([]models.TransactionResponse, int, any, int64, int, int)
}

type transactionRepo struct{}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepo{}
}

func (r *transactionRepo) Index(c *gin.Context) ([]models.TransactionResponse, int, any, int64, int, int) {
	var transactions []models.Transaction
	var total int64

	query := config.DB.Model(&models.Transaction{}).Preload("Mosque").Preload("QurbanPeriod").Preload("QurbanOffering")
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&transactions)

	var response []models.TransactionResponse
	for _, t := range transactions {
		response = append(response, models.TransactionResponse{
			ID:             t.ID,
			Code:           t.Code,
			CreatedAt:      t.CreatedAt,
			UpdatedAt:      t.UpdatedAt,
			Mosque:         t.Mosque,
			QurbanPeriod:   t.QurbanPeriod,
			QurbanOffering: t.QurbanOffering,
		})
	}

	return response, http.StatusOK, "transaction", total, page, limit
}

func (r *transactionRepo) IndexByMosqueID(c *gin.Context, mosqueID uint) ([]models.TransactionResponse, int, any, int64, int, int) {
	var transactions []models.Transaction
	var total int64

	query := config.DB.Model(&models.Transaction{}).Where("mosque_id = ?", mosqueID).Preload("Mosque").Preload("QurbanPeriod").Preload("QurbanOffering")
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&transactions)

	var response []models.TransactionResponse
	for _, t := range transactions {
		response = append(response, models.TransactionResponse{
			ID:             t.ID,
			Code:           t.Code,
			CreatedAt:      t.CreatedAt,
			UpdatedAt:      t.UpdatedAt,
			Mosque:         t.Mosque,
			QurbanPeriod:   t.QurbanPeriod,
			QurbanOffering: t.QurbanOffering,
		})
	}

	return response, http.StatusOK, "transaction", total, page, limit
}

func (r *transactionRepo) IndexByParticipantID(c *gin.Context, participantID uint) ([]models.TransactionResponse, int, any, int64, int, int) {
	var transactions []models.Transaction
	var total int64

	query := config.DB.Model(&models.Transaction{}).
		Joins("JOIN transaction_items ON transactions.id = transaction_items.transaction_id").
		Where("transaction_items.participant_id = ?", participantID).
		Preload("Mosque").Preload("QurbanPeriod").Preload("QurbanOffering")
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&transactions)

	var response []models.TransactionResponse
	for _, t := range transactions {
		response = append(response, models.TransactionResponse{
			ID:             t.ID,
			Code:           t.Code,
			CreatedAt:      t.CreatedAt,
			UpdatedAt:      t.UpdatedAt,
			Mosque:         t.Mosque,
			QurbanPeriod:   t.QurbanPeriod,
			QurbanOffering: t.QurbanOffering,
		})
	}

	return response, http.StatusOK, "transaction", total, page, limit
}

func (r *transactionRepo) Save(transaction *models.TransactionCreatePayload) (any, int, string, map[string]string) {
	tx := config.DB.Begin()

	var qurbanOffering models.QurbanOffering

	if err := config.DB.First(&qurbanOffering, transaction.QurbanOfferingID).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusNotFound, "qurban offering", nil
	}

	// Check if the qurban offering is still open
	if qurbanOffering.Status == models.Closed {
		tx.Rollback()
		return nil, http.StatusBadRequest, "transaction", map[string]string{"error": "Qurban offering is closed"}
	}

	// Validate mosque payment method if provided
	if transaction.MosquePaymentMethodID != nil {
		var paymentMethod models.MosquePaymentMethod
		if err := config.DB.First(&paymentMethod, transaction.MosquePaymentMethodID).Error; err != nil {
			tx.Rollback()
			return nil, http.StatusBadRequest, "transaction", map[string]string{"mosque_payment_method_id": "Payment method not found"}
		}

		// Ensure payment method belongs to the same mosque
		if paymentMethod.MosqueID != transaction.MosqueID {
			tx.Rollback()
			return nil, http.StatusBadRequest, "transaction", map[string]string{"mosque_payment_method_id": "Payment method does not belong to this mosque"}
		}

		// Ensure payment method is active
		if !paymentMethod.IsActive {
			tx.Rollback()
			return nil, http.StatusBadRequest, "transaction", map[string]string{"mosque_payment_method_id": "Payment method is not active"}
		}
	}

	// Validate at least one participant (existing or new)
	totalParticipants := len(transaction.ParticipantIDs) + len(transaction.Participants)
	if totalParticipants == 0 {
		tx.Rollback()
		return nil, http.StatusBadRequest, "transaction", map[string]string{"error": "At least one participant is required"}
	}

	// Check if offering has enough capacity for the participants
	availableSlots := qurbanOffering.Capacity - qurbanOffering.FilledSlots
	if totalParticipants > availableSlots {
		tx.Rollback()
		return nil, http.StatusBadRequest, "transaction", map[string]string{
			"error": fmt.Sprintf("Not enough slots available. Requested: %d, Available: %d", totalParticipants, availableSlots),
		}
	}

	// Collect all participant IDs (existing + newly created)
	var allParticipantIDs []uint

	// 1. Validate existing participants
	for _, participantID := range transaction.ParticipantIDs {
		var participant models.Participant
		if err := config.DB.First(&participant, participantID).Error; err != nil {
			tx.Rollback()
			return nil, http.StatusNotFound, "transaction", map[string]string{
				"participant_ids": "Participant ID " + fmt.Sprint(participantID) + " not found",
			}
		}
		allParticipantIDs = append(allParticipantIDs, participantID)
	}

	// 2. Create new participants
	for i, newParticipant := range transaction.Participants {
		// Get creator's participant data for address copying
		var creatorParticipant models.Participant
		var provinceCode, regencyCode, districtCode, villageCode *string
		var address *string

		if newParticipant.IsLikeAddress {
			if err := config.DB.Where("user_id = ?", transaction.CreatedByUserID).First(&creatorParticipant).Error; err == nil {
				provinceCode = creatorParticipant.ProvinceCode
				regencyCode = creatorParticipant.RegencyCode
				districtCode = creatorParticipant.DistrictCode
				villageCode = creatorParticipant.VillageCode
				address = creatorParticipant.Address
			}
		}

		participant := models.Participant{
			CreatedByUserID: transaction.CreatedByUserID,
			Name:            newParticipant.Name,
			Gender:          newParticipant.Gender,
			ProvinceCode:    provinceCode,
			RegencyCode:     regencyCode,
			DistrictCode:    districtCode,
			VillageCode:     villageCode,
			Address:         address,
		}

		if err := tx.Create(&participant).Error; err != nil {
			tx.Rollback()
			return nil, http.StatusInternalServerError, "transaction", map[string]string{
				"participants": "Failed to create participant at index " + fmt.Sprint(i),
			}
		}

		allParticipantIDs = append(allParticipantIDs, participant.ID)
	}

	// Calculate total amount
	// qurbanOffering.Price is per participant, so total = price × number of participants
	amountPerParticipant := qurbanOffering.Price
	totalAmount := qurbanOffering.Price * float64(totalParticipants)

	// Set expiration time (20 minutes from now)
	expiredAt := time.Now().Add(20 * time.Minute)

	// Create transaction with total amount
	transactionData := models.Transaction{
		MosqueID:              transaction.MosqueID,
		CreatedByUserID:       transaction.CreatedByUserID,
		PaidAmount:            &totalAmount,
		QurbanPeriodID:        qurbanOffering.QurbanPeriodID,
		QurbanOfferingID:      transaction.QurbanOfferingID,
		MosquePaymentMethodID: transaction.MosquePaymentMethodID,
		PaymentNote:           transaction.PaymentNote,
		ExpiredAt:             &expiredAt,
	}

	if err := tx.Create(&transactionData).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "transaction", nil
	}

	// Create transaction items for all participants
	for _, participantID := range allParticipantIDs {
		transactionItem := models.TransactionItem{
			TransactionID: transactionData.ID,
			ParticipantID: participantID,
			Amount:        &amountPerParticipant,
		}

		if err := tx.Create(&transactionItem).Error; err != nil {
			tx.Rollback()
			return nil, http.StatusInternalServerError, "transaction item", nil
		}
	}

	// Update qurban offering filled slots
	qurbanOffering.FilledSlots += totalParticipants
	if err := tx.Save(&qurbanOffering).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "qurban offering", nil
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "transaction", nil
	}

	return transactionData, http.StatusCreated, "transaction", nil
}

func (r *transactionRepo) UpdateProof(payload *models.TransactionUploadProofPayload) (any, int, string, map[string]string) {
	var transaction models.Transaction

	if err := config.DB.First(&transaction, payload.ID).Error; err != nil {
		return nil, http.StatusNotFound, "transaction", map[string]string{
			"error": "Transaction not found",
		}
	}

	// Update transaction with proof details
	transaction.ProofURL = &payload.ProofURL
	transaction.PaidAmount = &payload.PaidAmount
	transaction.PaymentNote = payload.PaymentNote
	transaction.PaymentStatus = models.TransactionWaitingVerification

	if err := config.DB.Save(&transaction).Error; err != nil {
		return nil, http.StatusInternalServerError, "transaction", map[string]string{
			"error": "Failed to update transaction proof",
		}
	}

	return transaction, http.StatusOK, "transaction", nil
}
