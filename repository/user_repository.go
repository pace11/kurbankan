package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.UserListResponse, int, any, int64, int, int)
	Update(id uint, user *models.User) (any, int, string, map[string]string)
}

type userRepo struct{}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (r *userRepo) Index(c *gin.Context, filters map[string]any) ([]models.UserListResponse, int, any, int64, int, int) {
	var users []models.User
	var total int64

	query := utils.FilterByParams(config.DB.Model(&models.User{}), filters)
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&users)

	var response []models.UserListResponse
	for _, u := range users {
		response = append(response, models.UserListResponse{
			ID:        u.ID,
			Email:     u.Email,
			Role:      u.Role,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}

	return response, http.StatusOK, "user", total, page, limit
}

func (r *userRepo) Update(id uint, user *models.User) (any, int, string, map[string]string) {
	var existing models.User

	if err := config.DB.First(&existing, id).Error; err != nil {
		return nil, http.StatusNotFound, "user", nil
	}

	if user.Password != "" {
		hashed, err := utils.HashPassword(user.Password)
		if err != nil {
			return nil, http.StatusInternalServerError, "user", nil
		}
		existing.Password = hashed
	}

	if err := config.DB.Model(&existing).Updates(map[string]any{
		"email": user.Email,
		"role":  user.Role,
	}).Error; err != nil {
		return nil, http.StatusInternalServerError, "user", nil
	}

	return user, http.StatusOK, "user", nil
}
