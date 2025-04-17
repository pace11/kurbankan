package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
)

type UserRepository interface {
	Index() []models.User
	Save(user *models.User)
	Update(id uint, user *models.User) bool
	Delete(id uint) bool
}

type userRepo struct{}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (r *userRepo) Index() []models.User {
	var user []models.User
	config.DB.Find(&user)
	return user
}

func (r *userRepo) Save(user *models.User) {
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return
	}
	user.Password = hashed
	config.DB.Create(user)
}

func (r *userRepo) Update(id uint, user *models.User) bool {
	var existing models.User
	if err := config.DB.First(&existing, id).Error; err != nil {
		return false
	}
	existing.Email = user.Email
	existing.Role = user.Role

	if user.Password != "" {
		hashed, err := utils.HashPassword(user.Password)
		if err != nil {
			return false
		}
		existing.Password = hashed
	}

	config.DB.Save(&existing)
	return true
}

func (r *userRepo) Delete(id uint) bool {
	result := config.DB.Delete(&models.User{}, id)
	return result.RowsAffected > 0
}
