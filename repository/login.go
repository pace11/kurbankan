package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"
)

type LoginRepository interface {
	Login(payload *models.LoginPayload) (any, int, string, map[string]string)
}

type loginRepository struct{}

func NewLoginRepository() LoginRepository {
	return &loginRepository{}
}

func (r *loginRepository) Login(payload *models.LoginPayload) (any, int, string, map[string]string) {
	var user models.User

	if err := config.DB.Where("email = ?", payload.Email).First(&user).Error; err != nil {
		return nil, http.StatusNotFound, "user", nil
	}

	if !utils.CheckPasswordHash(payload.Password, user.Password) {
		return nil, http.StatusUnauthorized, "invalid credentials", nil
	}

	token, err := utils.GenerateToken(user.ID, user.Email, user.PlatformRole)
	if err != nil {
		return nil, http.StatusInternalServerError, "generate token", nil
	}

	userResponse := models.UserResponse{
		ID:           user.ID,
		Email:        user.Email,
		PlatformRole: user.PlatformRole,
		CreatedAt:    &user.CreatedAt,
		UpdatedAt:    &user.UpdatedAt,
	}

	response := map[string]any{
		"token": token,
		"user":  userResponse,
	}

	return response, http.StatusOK, "user", nil
}
