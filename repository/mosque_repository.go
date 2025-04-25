package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MosqueRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.MosqueResponse, int, any, int64, int, int)
	Show(id uint) (*models.MosqueResponse, int, string, map[string]string)
	Save(mosque *models.UserCreateDTO) (any, int, string, map[string]string)
	Update(id uint, mosque *models.UserUpdateDTO) (any, int, string, map[string]string)
	Delete(id uint) (any, int, string, map[string]string)
}

type mosqueRepository struct{}

func NewMosqueRepository() MosqueRepository {
	return &mosqueRepository{}
}

func (r *mosqueRepository) Index(c *gin.Context, filters map[string]any) ([]models.MosqueResponse, int, any, int64, int, int) {
	var mosques []models.Mosque
	var total int64

	query := utils.FilterByParams(config.DB.Model(&models.Mosque{}).Preload("Province").Preload("Regency").Preload("District").Preload("Village").Preload("User"), filters)
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&mosques)

	var response []models.MosqueResponse
	for _, m := range mosques {
		response = append(response, models.MosqueResponse{
			ID:        m.ID,
			Name:      m.Name,
			Address:   m.Address,
			Photos:    m.Photos,
			Province:  m.Province,
			Regency:   m.Regency,
			District:  m.District,
			Village:   m.Village,
			User:      models.ToUserResponse(m.User),
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		})
	}

	return response, http.StatusOK, "mosque", total, page, limit
}

func (r *mosqueRepository) Show(id uint) (*models.MosqueResponse, int, string, map[string]string) {
	var mosque models.Mosque
	err := config.DB.Preload("Province").Preload("Regency").Preload("District").Preload("Village").Preload("User").Where("id = ?", id).First(&mosque).Error

	if err != nil {
		return nil, http.StatusInternalServerError, "mosque", nil
	}

	response := &models.MosqueResponse{
		ID:        mosque.ID,
		Name:      mosque.Name,
		Address:   mosque.Address,
		Photos:    mosque.Photos,
		Province:  mosque.Province,
		Regency:   mosque.Regency,
		District:  mosque.District,
		Village:   mosque.Village,
		User:      models.ToUserResponse(mosque.User),
		CreatedAt: mosque.CreatedAt,
		UpdatedAt: mosque.UpdatedAt,
	}
	return response, http.StatusOK, "mosque", nil
}

func (r *mosqueRepository) Save(mosque *models.UserCreateDTO) (any, int, string, map[string]string) {
	tx := config.DB.Begin()

	hashed, err := utils.HashPassword(mosque.Password)
	if err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "mosque", nil
	}

	userToCreate := models.User{
		Email:    mosque.Email,
		Password: hashed,
		Role:     models.MosqueMember,
	}

	if err := tx.Save(&userToCreate).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "user mosque", nil
	}

	mosqueCreate := models.Mosque{
		UserID:       userToCreate.ID,
		Name:         mosque.Name,
		Address:      mosque.Address,
		ProvinceCode: mosque.ProvinceCode,
		RegencyCode:  mosque.RegencyCode,
		DistrictCode: mosque.DistrictCode,
		VillageCode:  mosque.VillageCode,
	}

	if err := tx.Save(&mosqueCreate).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "mosque", nil
	}

	return mosque, http.StatusCreated, "mosque", nil
}

func (r *mosqueRepository) Update(id uint, mosque *models.UserUpdateDTO) (any, int, string, map[string]string) {
	var existing models.Mosque
	var userToUpdate models.User

	if err := config.DB.Preload("User").First(&existing, id).Error; err != nil {
		return nil, http.StatusNotFound, "mosque", nil
	}

	tx := config.DB.Begin()

	if mosque.Password != "" {
		hashed, err := utils.HashPassword(mosque.Password)
		if err != nil {
			tx.Rollback()
			return nil, http.StatusInternalServerError, "mosque", nil
		}
		mosque.Password = hashed
	} else {
		mosque.Password = existing.User.Password
	}

	if err := tx.First(&userToUpdate, existing.UserID).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusNotFound, "user mosque", nil
	}

	if err := tx.Model(&userToUpdate).Updates(map[string]any{
		"email":    mosque.Email,
		"password": mosque.Password,
		"role":     models.UserRole(*mosque.Role),
	}).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "mosque", nil
	}

	if err := tx.Save(&existing).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "mosque", nil
	}

	if err := tx.Model(&existing).Updates(map[string]any{
		"name":          mosque.Name,
		"address":       mosque.Address,
		"photos":        mosque.Photos,
		"province_code": mosque.ProvinceCode,
		"district_code": mosque.DistrictCode,
		"regency_code":  mosque.RegencyCode,
		"village_code":  mosque.VillageCode,
	}).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "mosque", nil
	}

	if err := tx.Commit().Error; err != nil {
		return nil, http.StatusInternalServerError, "mosque", nil
	}

	return mosque, http.StatusOK, "mosque", nil
}

func (r *mosqueRepository) Delete(id uint) (any, int, string, map[string]string) {
	result := config.DB.Delete(&models.Mosque{}, id)

	if result.Error != nil {
		return nil, http.StatusInternalServerError, "mosque", nil
	}

	if result.RowsAffected == 0 {
		return nil, http.StatusNotFound, "mosque", nil
	}

	return nil, http.StatusOK, "mosque", nil
}
