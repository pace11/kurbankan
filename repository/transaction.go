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
	Create(transaction *models.TransactionCreatePayload) (any, int, string, map[string]string)
	UpdateProof(payload *models.TransactionUploadProofPayload) (any, int, string, map[string]string)
	Verify(payload *models.TransactionVerifyPayload) (any, int, string, map[string]string)
	MarkExpiredPendingTransactions() (any, int, string, map[string]string)

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

func (r *transactionRepo) Create(transaction *models.TransactionCreatePayload) (any, int, string, map[string]string) {
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

func (r *transactionRepo) Verify(payload *models.TransactionVerifyPayload) (any, int, string, map[string]string) {
	// Start a database transaction
	tx := config.DB.Begin()

	var transaction models.Transaction
	if err := tx.Preload("QurbanOffering").First(&transaction, payload.ID).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusNotFound, "transaction", map[string]string{"error": "Transaction not found"}
	}

	// Validate transaction status
	if transaction.PaymentStatus != models.TransactionWaitingVerification {
		tx.Rollback()
		return nil, http.StatusBadRequest, "transaction", map[string]string{"error": "Transaction is not in waiting_verification status"}
	}

	// Validate rejected reason if status is rejected
	if payload.PaymentStatus == models.TransactionRejected && (payload.RejectedReason == nil || *payload.RejectedReason == "") {
		tx.Rollback()
		return nil, http.StatusBadRequest, "transaction", map[string]string{"rejected_reason": "Rejected reason is required when rejecting a transaction"}
	}

	// Count transaction items (participants) for this transaction
	var transactionItemCount int64
	if err := tx.Model(&models.TransactionItem{}).Where("transaction_id = ?", payload.ID).Count(&transactionItemCount).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "transaction", map[string]string{"error": "Failed to count transaction items"}
	}

	// Update transaction status
	now := time.Now()
	transaction.PaymentStatus = payload.PaymentStatus
	transaction.VerifiedByUserID = &payload.VerifiedByUserID
	transaction.VerifiedAt = &now
	transaction.RejectedReason = payload.RejectedReason

	switch payload.PaymentStatus {
	case models.TransactionPaid:
		transaction.PaidAt = &now

		// Update ConfirmedSlots in QurbanOffering
		if transaction.QurbanOffering != nil {
			transaction.QurbanOffering.ConfirmedSlots += int(transactionItemCount)

			// Auto-close offering when all slots are confirmed/paid
			if transaction.QurbanOffering.ConfirmedSlots >= transaction.QurbanOffering.Capacity {
				transaction.QurbanOffering.Status = models.Closed
			}

			if err := tx.Save(transaction.QurbanOffering).Error; err != nil {
				tx.Rollback()
				return nil, http.StatusInternalServerError, "transaction", map[string]string{"error": "Failed to update confirmed slots"}
			}
		}
	case models.TransactionRejected:
		// Decrease FilledSlots when transaction is rejected
		if transaction.QurbanOffering != nil {
			transaction.QurbanOffering.FilledSlots -= int(transactionItemCount)
			if transaction.QurbanOffering.FilledSlots < 0 {
				transaction.QurbanOffering.FilledSlots = 0
			}

			// Re-open offering if it was closed and now has available slots
			if transaction.QurbanOffering.Status == models.Closed &&
				transaction.QurbanOffering.ConfirmedSlots < transaction.QurbanOffering.Capacity {
				transaction.QurbanOffering.Status = models.Open
			}

			if err := tx.Save(transaction.QurbanOffering).Error; err != nil {
				tx.Rollback()
				return nil, http.StatusInternalServerError, "transaction", map[string]string{"error": "Failed to update filled slots"}
			}
		}
	}

	if err := tx.Save(&transaction).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "transaction", map[string]string{"error": "Failed to verify transaction"}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "transaction", map[string]string{"error": "Failed to commit verification"}
	}

	return nil, http.StatusAccepted, "transaction", nil
}

func (r *transactionRepo) MarkExpiredPendingTransactions() (any, int, string, map[string]string) {
	// Start a database transaction
	tx := config.DB.Begin()

	var transactions []models.Transaction
	now := time.Now()

	// Find all pending transactions that have expired
	if err := tx.Preload("QurbanOffering").Where("payment_status = ? AND expired_at IS NOT NULL AND expired_at < ?", models.TransactionPending, now).Find(&transactions).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "transaction", map[string]string{"error": "Failed to fetch expired transactions"}
	}

	if len(transactions) == 0 {
		tx.Rollback()
		return nil, http.StatusOK, "transaction", nil
	}

	cancelledIDs := []uint{}
	offeringUpdates := make(map[uint]*models.QurbanOffering) // Track offerings to update

	for _, transaction := range transactions {
		// Count transaction items (participants) for this transaction
		var transactionItemCount int64
		if err := tx.Model(&models.TransactionItem{}).Where("transaction_id = ?", transaction.ID).Count(&transactionItemCount).Error; err != nil {
			continue // Skip this transaction on error
		}

		// Update transaction status to cancelled
		transaction.PaymentStatus = models.TransactionCancelled
		if err := tx.Save(&transaction).Error; err != nil {
			continue // Skip this transaction on error
		}

		cancelledIDs = append(cancelledIDs, transaction.ID)

		// Decrease FilledSlots from QurbanOffering
		if transaction.QurbanOffering != nil {
			offeringID := transaction.QurbanOffering.ID

			// Get or initialize offering in map
			if _, exists := offeringUpdates[offeringID]; !exists {
				offeringUpdates[offeringID] = transaction.QurbanOffering
				offeringUpdates[offeringID].FilledSlots = transaction.QurbanOffering.FilledSlots
			}

			// Decrease filled slots
			offeringUpdates[offeringID].FilledSlots -= int(transactionItemCount)
			if offeringUpdates[offeringID].FilledSlots < 0 {
				offeringUpdates[offeringID].FilledSlots = 0
			}
		}
	}

	// Update all affected offerings
	for _, offering := range offeringUpdates {
		// Re-open offering if it was closed and now has available slots
		if offering.Status == models.Closed && offering.ConfirmedSlots < offering.Capacity {
			offering.Status = models.Open
		}

		if err := tx.Save(offering).Error; err != nil {
			tx.Rollback()
			return nil, http.StatusInternalServerError, "transaction", map[string]string{"error": "Failed to update qurban offering slots"}
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "transaction", map[string]string{"error": "Failed to commit transaction"}
	}

	return nil, http.StatusAccepted, "mark expired transactions", nil
}
