package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"
)

type LoginRepository interface {
	Login(login *models.LoginDTO) (any, int, string, map[string]string)
}

type loginRepository struct{}

func NewLoginRepository() LoginRepository {
	return &loginRepository{}
}

func (r *loginRepository) Login(login *models.LoginDTO) (any, int, string, map[string]string) {
	var user models.User

	if err := config.DB.Where("email = ?", login.Email).First(&user).Error; err != nil {
		return nil, http.StatusNotFound, "user", nil
	}

	if !utils.CheckPasswordHash(login.Password, user.Password) {
		return nil, http.StatusUnauthorized, "invalid credentials", nil
	}

	token, err := utils.GenerateToken(user.ID, user.Email, string(user.Role))
	if err != nil {
		return nil, http.StatusInternalServerError, "generate token", nil
	}

	userResponse := models.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  string(user.Role),
	}

	response := map[string]any{
		"token": token,
		"user":  userResponse,
	}

	return response, http.StatusOK, "user", nil
}
