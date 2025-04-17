package controllers

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterMosqueDTO struct {
	Email        string  `json:"email" binding:"required,email"`
	Password     string  `json:"password" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	Address      *string `json:"address"`
	ProvinceCode string  `json:"province_code" binding:"required"`
	RegencyCode  string  `json:"regency_code" binding:"required"`
	DistrictCode string  `json:"district_code" binding:"required"`
	VillageCode  string  `json:"village_code" binding:"required"`
}

func RegisterMosque(c *gin.Context) {
	var payload RegisterMosqueDTO
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
		Role:     models.UserMember,
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
