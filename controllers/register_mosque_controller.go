package controllers

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterMosque(c *gin.Context) {
	var payload models.UserCreateDTO
	if utils.BindAndValidate(c, &payload) != nil {
		return
	}

	tx := config.DB.Begin()

	hashed, err := utils.HashPassword(payload.Password)
	if err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := models.User{
		Email:    payload.Email,
		Password: hashed,
		Role:     models.MosqueMember,
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusBadRequest, "Email already exists or invalid mosque data")
		return
	}

	mosque := models.Mosque{
		UserID:       user.ID,
		Name:         payload.Name,
		Address:      payload.Address,
		Photos:       payload.Photos,
		ProvinceCode: payload.ProvinceCode,
		RegencyCode:  payload.RegencyCode,
		DistrictCode: payload.DistrictCode,
		VillageCode:  payload.VillageCode,
	}

	if err := tx.Create(&mosque).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to create mosque")
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Mosque registered successfully",
		"data": gin.H{
			"user_id":   user.ID,
			"mosque_id": mosque.ID,
			"email":     user.Email,
			"name":      mosque.Name,
			"address":   mosque.Address,
		},
	})
}
