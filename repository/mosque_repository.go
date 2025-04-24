package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MosqueRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.MosqueResponse, int, any, any, int64, int, int)
	Show(id uint) (*models.MosqueResponse, error)
	Save(mosque *models.UserCreateDTO) bool
	Update(id uint, mosque *models.UserUpdateDTO) bool
	Delete(id uint) bool
}

type mosqueRepository struct{}

func NewMosqueRepository() MosqueRepository {
	return &mosqueRepository{}
}

func (r *mosqueRepository) Index(c *gin.Context, filters map[string]any) ([]models.MosqueResponse, int, any, any, int64, int, int) {
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

	return response, http.StatusOK, "mosque", "get", total, page, limit
}

func (r *mosqueRepository) Show(id uint) (*models.MosqueResponse, error) {
	var mosque models.Mosque
	err := config.DB.Preload("Province").Preload("Regency").Preload("District").Preload("Village").Preload("User").Where("id = ?", id).First(&mosque).Error

	if err != nil {
		return nil, err
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
	return response, nil
}

func (r *mosqueRepository) Save(mosque *models.UserCreateDTO) bool {
	tx := config.DB.Begin()

	hashed, err := utils.HashPassword(mosque.Password)
	if err != nil {
		tx.Rollback()
		return false
	}

	userToCreate := models.User{
		Email:    mosque.Email,
		Password: hashed,
		Role:     models.MosqueMember,
	}

	if err := tx.Save(&userToCreate).Error; err != nil {
		tx.Rollback()
		return false
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
		return false
	}

	return tx.Commit().Error == nil
}

func (r *mosqueRepository) Update(id uint, mosque *models.UserUpdateDTO) bool {
	var existing models.Mosque
	var userToUpdate models.User

	if err := config.DB.Preload("User").First(&existing, id).Error; err != nil {
		return false
	}

	tx := config.DB.Begin()

	if mosque.Password != "" {
		hashed, err := utils.HashPassword(mosque.Password)
		if err != nil {
			tx.Rollback()
			return false
		}
		mosque.Password = hashed
	} else {
		mosque.Password = existing.User.Password
	}

	if err := tx.First(&userToUpdate, existing.UserID).Error; err != nil {
		tx.Rollback()
		return false
	}

	userToUpdate.Email = mosque.Email
	userToUpdate.Password = mosque.Password
	userToUpdate.Role = models.UserRole(*mosque.Role)

	if err := tx.Save(&userToUpdate).Error; err != nil {
		tx.Rollback()
		return false
	}

	existing.Name = mosque.Name
	existing.Address = mosque.Address
	existing.Photos = mosque.Photos
	existing.ProvinceCode = mosque.ProvinceCode
	existing.DistrictCode = mosque.DistrictCode
	existing.RegencyCode = mosque.RegencyCode
	existing.VillageCode = mosque.VillageCode

	if err := tx.Save(&existing).Error; err != nil {
		tx.Rollback()
		return false
	}

	return tx.Commit().Error == nil
}

func (r *mosqueRepository) Delete(id uint) bool {
	result := config.DB.Delete(&models.Mosque{}, id)
	return result.RowsAffected > 0
}
