package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"
)

type RegisterRepository interface {
	SaveParticipant(participant *models.UserCreateDTO) (any, int, string, map[string]string)
	SaveMosque(mosque *models.UserCreateDTO) (any, int, string, map[string]string)
}

type registerRepository struct{}

func NewRegisterRepository() RegisterRepository {
	return &registerRepository{}
}

func (r *registerRepository) SaveParticipant(participant *models.UserCreateDTO) (any, int, string, map[string]string) {
	tx := config.DB.Begin()

	hashed, err := utils.HashPassword(participant.Password)
	if err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "participant", nil
	}

	userToCreate := models.User{
		Email:    participant.Email,
		Password: hashed,
		Role:     models.UserMember,
	}

	if err := tx.Save(&userToCreate).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "participant", nil
	}

	participantCreate := models.Participant{
		UserID:       userToCreate.ID,
		Name:         participant.Name,
		Address:      participant.Address,
		ProvinceCode: participant.ProvinceCode,
		RegencyCode:  participant.RegencyCode,
		DistrictCode: participant.DistrictCode,
		VillageCode:  participant.VillageCode,
	}

	if err := tx.Save(&participantCreate).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "participant", nil
	}

	if tx.Commit().Error != nil {
		return nil, http.StatusInternalServerError, "trx participant", nil
	}

	return participant, http.StatusCreated, "participant", nil
}

func (r *registerRepository) SaveMosque(mosque *models.UserCreateDTO) (any, int, string, map[string]string) {
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
		return nil, http.StatusInternalServerError, "mosque", nil
	}

	mosqueCreate := models.Mosque{
		UserID:       userToCreate.ID,
		Name:         mosque.Name,
		Address:      mosque.Address,
		Photos:       mosque.Photos,
		ProvinceCode: mosque.ProvinceCode,
		RegencyCode:  mosque.RegencyCode,
		DistrictCode: mosque.DistrictCode,
		VillageCode:  mosque.VillageCode,
	}

	if err := tx.Save(&mosqueCreate).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, "mosque", nil
	}

	if tx.Commit().Error != nil {
		return nil, http.StatusInternalServerError, "trx mosque", nil
	}

	return mosque, http.StatusCreated, "mosque", nil
}
