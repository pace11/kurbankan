package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ParticipantRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.ParticipantResponse, int, any, int64, int, int)
	Show(id uint) (*models.ParticipantResponse, int, string, map[string]string)
	Update(id uint, participant *models.UserUpdateDTO) (any, int, string, map[string]string)
	Delete(id uint) (any, int, string, map[string]string)
}

type participantRepository struct{}

func NewParticipantRepository() ParticipantRepository {
	return &participantRepository{}
}

func (r *participantRepository) Index(c *gin.Context, filters map[string]any) ([]models.ParticipantResponse, int, any, int64, int, int) {
	var participants []models.Participant
	var total int64

	query := utils.FilterByParams(config.DB.Model(&models.Participant{}).Preload("Province").Preload("Regency").Preload("District").Preload("Village").Preload("User"), filters)
	query.Count(&total)

	paginatedQuery, page, limit := utils.ApplyPagination(c, query)
	paginatedQuery.Find(&participants)

	var response []models.ParticipantResponse
	for _, p := range participants {
		response = append(response, models.ParticipantResponse{
			ID:        p.ID,
			Name:      p.Name,
			Address:   p.Address,
			Province:  p.Province,
			Regency:   p.Regency,
			District:  p.District,
			Village:   p.Village,
			User:      models.ToUserResponse(p.User),
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	return response, http.StatusOK, "participant", total, page, limit
}

func (r *participantRepository) Show(id uint) (*models.ParticipantResponse, int, string, map[string]string) {
	var participant models.Participant

	if err := config.DB.Preload("Province").Preload("Regency").Preload("District").Preload("Village").Preload("User").Where("id = ?", id).First(&participant).Error; err != nil {
		return nil, http.StatusNotFound, "participant", nil
	}

	response := &models.ParticipantResponse{
		ID:        participant.ID,
		Name:      participant.Name,
		Address:   participant.Address,
		Province:  participant.Province,
		Regency:   participant.Regency,
		District:  participant.District,
		Village:   participant.Village,
		User:      models.ToUserResponse(participant.User),
		CreatedAt: participant.CreatedAt,
		UpdatedAt: participant.UpdatedAt,
	}

	return response, http.StatusOK, "participant", nil
}

func (r *participantRepository) Update(id uint, participant *models.UserUpdateDTO) (any, int, string, map[string]string) {
	var existing models.Participant
	var userToUpdate models.User

	if err := config.DB.Preload("User").First(&existing, id).Error; err != nil {
		return nil, http.StatusNotFound, "participant", nil
	}

	tx := config.DB.Begin()

	if participant.Password != "" {
		hashed, err := utils.HashPassword(participant.Password)
		if err != nil {
			tx.Rollback()
			return nil, http.StatusInternalServerError, "participant", nil
		}
		participant.Password = hashed
	} else {
		participant.Password = existing.User.Password
	}

	if err := tx.First(&userToUpdate, existing.UserID).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusNotFound, "user participant", nil
	}

	if err := tx.Model(&userToUpdate).Updates(map[string]any{
		"email":    participant.Email,
		"password": participant.Password,
		"role":     models.UserRole(*participant.Role),
	}).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "participant", nil
	}

	if err := tx.Model(&existing).Updates(map[string]any{
		"name":          participant.Name,
		"address":       participant.Address,
		"province_code": participant.ProvinceCode,
		"district_code": participant.DistrictCode,
		"regency_code":  participant.RegencyCode,
		"village_code":  participant.VillageCode,
	}).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "participant", nil
	}

	if err := tx.Commit().Error; err != nil {
		return nil, http.StatusInternalServerError, "participant", nil
	}

	return participant, http.StatusOK, "participant", nil
}

func (r *participantRepository) Delete(id uint) (any, int, string, map[string]string) {
	result := config.DB.Delete(&models.Participant{}, id)

	if result.Error != nil {
		return nil, http.StatusInternalServerError, "participant", nil
	}

	if result.RowsAffected == 0 {
		return nil, http.StatusNotFound, "mosque", nil
	}

	return nil, http.StatusOK, "mosque", nil
}
