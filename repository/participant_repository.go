package repository

import (
	"kurbankan/config"
	"kurbankan/models"
)

type ParticipantRepository interface {
	Index() []models.ParticipantResponse
	Show(id uint) (*models.ParticipantResponse, error)
}

type participantRepository struct{}

func NewParticipantRepository() ParticipantRepository {
	return &participantRepository{}
}

func (r *participantRepository) Index() []models.ParticipantResponse {
	var participants []models.Participant
	config.DB.Preload("Province").Preload("Regency").Preload("District").Preload("Village").Preload("User").Find(&participants)

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
	return response
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
