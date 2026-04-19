package utils

import (
	"errors"
	"kurbankan/config"
	"kurbankan/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMosqueMemberByUserID retrieves mosque data for a given user ID
func GetMosqueMemberByUserID(ctx *gin.Context, userID uint, role string) (*models.MosqueMember, error, int, string, map[string]string) {
	var mosqueMember models.MosqueMember

	if err := config.DB.Where("user_id = ?", userID).First(&mosqueMember).Error; err != nil {
		return nil, errors.New("mosque not found for the user"), http.StatusNotFound, "mosque", map[string]string{"error": "mosque not found for the user"}
	}

	return &mosqueMember, nil, http.StatusOK, "mosque", nil
}

// GetParticipantByUserID retrieves participant data for a given user ID
func GetParticipantByUserID(ctx *gin.Context, userID uint, role string) (*models.Participant, error, int, string, map[string]string) {
	var participant models.Participant

	if err := config.DB.Where("user_id = ?", userID).First(&participant).Error; err != nil {
		return nil, errors.New("participant not found for the user"), http.StatusNotFound, "participant", map[string]string{"error": "participant not found for the user"}
	}

	return &participant, nil, http.StatusOK, "participant", nil
}
