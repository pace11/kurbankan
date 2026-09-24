package repository

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"
)

type RegisterRepository interface {
	SaveParticipant(participant *models.UserCreatePayload) (any, int, map[string]string)
	SaveMosque(mosque *models.UserCreatePayload) (any, int, map[string]string)
}

type registerRepository struct{}

func NewRegisterRepository() RegisterRepository {
	return &registerRepository{}
}

func (r *registerRepository) SaveParticipant(participant *models.UserCreatePayload) (any, int, map[string]string) {
	tx := config.DB.Begin()

	hashed, err := utils.HashPassword(participant.Password)
	if err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, nil
	}

	// Create user first
	platformRole := models.PlatformRoleParticipant
	userToCreate := models.User{
		Email:        participant.Email,
		Password:     hashed,
		PlatformRole: &platformRole,
	}

	if err := tx.Save(&userToCreate).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, nil
	}

	participantCreate := models.Participant{
		UserID:          &userToCreate.ID,
		CreatedByUserID: userToCreate.ID,
		Name:            participant.Name,
		Gender:          participant.Gender,
		Address:         participant.Address,
		ProvinceCode:    &participant.ProvinceCode,
		RegencyCode:     &participant.RegencyCode,
		DistrictCode:    &participant.DistrictCode,
		VillageCode:     &participant.VillageCode,
	}

	if err := tx.Save(&participantCreate).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, nil
	}

	if tx.Commit().Error != nil {
		return nil, http.StatusInternalServerError, nil
	}

	return participant, http.StatusCreated, nil
}

func (r *registerRepository) SaveMosque(mosque *models.UserCreatePayload) (any, int, map[string]string) {
	tx := config.DB.Begin()

	hashed, err := utils.HashPassword(mosque.Password)
	if err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, nil
	}

	// Create user first
	platformRole := models.PlatformRoleMosque
	userToCreate := models.User{
		Email:        mosque.Email,
		Password:     hashed,
		PlatformRole: &platformRole,
	}

	if err := tx.Save(&userToCreate).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, nil
	}

	// Then create mosque with the created user's ID
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
		return nil, http.StatusInternalServerError, nil
	}

	// Create mosque member with the created user's ID and mosque's ID, and set role as admin
	mosqueMemberCreate := models.MosqueMember{
		MosqueID: mosqueCreate.ID,
		UserID:   userToCreate.ID,
		Role:     models.MosqueAdmin,
	}

	if err := tx.Save(&mosqueMemberCreate).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, nil
	}

	if tx.Commit().Error != nil {
		return nil, http.StatusInternalServerError, nil
	}

	return mosque, http.StatusCreated, nil
}
