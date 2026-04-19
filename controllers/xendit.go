package controllers

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type XenditVAWebhookPayload struct {
	ExternalID string  `json:"external_id"`
	Amount     float64 `json:"amount"`
}

// Handle webhook callback from Xendit VA payment
func XenditVAWebhookHandler(c *gin.Context) {
	var payload XenditVAWebhookPayload
	var transactionItems models.TransactionItem
	if utils.BindAndValidate(c, &payload) != nil {
		return
	}

	// fmt.Printf("Received Xendit VA Webhook: %+v\n", payload)

	if err := config.DB.Where("external_id = ?", payload.ExternalID).First(&transactionItems).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Transaction detail not found")
		return
	}

	// if payload.Status == "COMPLETED" {
	// 	// Update status to "paid" if match found
	// 	result := config.DB.Model(&models.TransactionItem{}).
	// 		Where("external_id = ?", payload.ExternalID).
	// 		UpdateColumns(map[string]interface{}{
	// 			"status":     models.Paid,
	// 			"paid_at":    time.Now(),
	// 			"updated_at": time.Now(),
	// 		})

	// 	if result.RowsAffected > 0 {
	// 		c.JSON(http.StatusOK, gin.H{"message": "Transaction updated"})
	// 	} else {
	// 		c.JSON(http.StatusNotFound, gin.H{"message": "Transaction not found"})
	// 	}
	// 	return
	// }

	// Handle other status if needed
	c.JSON(http.StatusOK, gin.H{"data": transactionItems})
}
