package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ParticipantRepository interface {
	Index(c *gin.Context, filters map[string]any) ([]models.ParticipantResponse, int, any, any, int64, int, int)
	Show(id uint) (*models.ParticipantResponse, error)
	Update(id uint, participant *models.UserUpdateDTO) bool
	Delete(id uint) bool
}

type participantRepository struct{}

func NewParticipantRepository() ParticipantRepository {
	return &participantRepository{}
}

func (r *participantRepository) Index(c *gin.Context, filters map[string]any) ([]models.ParticipantResponse, int, any, any, int64, int, int) {
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
	return response, http.StatusOK, "participant", "get", total, page, limit
}

func (r *participantRepository) Show(id uint) (*models.ParticipantResponse, error) {
	var participant models.Participant
	err := config.DB.Preload("Province").Preload("Regency").Preload("District").Preload("Village").Preload("User").Where("id = ?", id).First(&participant).Error

	if err != nil {
		return nil, err
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
	return response, nil
}

func (r *participantRepository) Update(id uint, participant *models.UserUpdateDTO) bool {
	var existing models.Participant
	var userToUpdate models.User

	if err := config.DB.Preload("User").First(&existing, id).Error; err != nil {
		return false
	}

	tx := config.DB.Begin()

	if participant.Password != "" {
		hashed, err := utils.HashPassword(participant.Password)
		if err != nil {
			tx.Rollback()
			return false
		}
		participant.Password = hashed
	} else {
		participant.Password = existing.User.Password
	}

	if err := tx.First(&userToUpdate, existing.UserID).Error; err != nil {
		tx.Rollback()
		return false
	}

	userToUpdate.Email = participant.Email
	userToUpdate.Password = participant.Password
	userToUpdate.Role = models.UserRole(*participant.Role)

	if err := tx.Save(&userToUpdate).Error; err != nil {
		tx.Rollback()
		return false
	}

	existing.Name = participant.Name
	existing.Address = participant.Address
	existing.ProvinceCode = participant.ProvinceCode
	existing.DistrictCode = participant.DistrictCode
	existing.RegencyCode = participant.RegencyCode
	existing.VillageCode = participant.VillageCode

	if err := tx.Save(&existing).Error; err != nil {
		tx.Rollback()
		return false
	}

	return tx.Commit().Error == nil
}

func (r *participantRepository) Delete(id uint) bool {
	result := config.DB.Delete(&models.Participant{}, id)
	return result.RowsAffected > 0
}
