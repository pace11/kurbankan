package repository

import (
	"kurbankan/config"
	"kurbankan/models"
)

type MosqueRepository interface {
	Index() []models.MosqueResponse
	Show(id uint) (*models.MosqueResponse, error)
}

type mosqueRepository struct{}

func NewMosqueRepository() MosqueRepository {
	return &mosqueRepository{}
}

func (r *mosqueRepository) Index() []models.MosqueResponse {
	var mosques []models.Mosque
	config.DB.Preload("Province").Preload("Regency").Preload("District").Preload("Village").Preload("User").Find(&mosques)

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
	return response
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
