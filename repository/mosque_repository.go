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
	config.DB.Preload("Province").Preload("Regency").Preload("District").Preload("Village").Find(&mosques)

	var response []models.MosqueResponse
	for _, m := range mosques {
		response = append(response, models.MosqueResponse{
			ID:        m.ID,
			Name:      m.Name,
			Address:   m.Address,
			Photos:    m.Photos,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
			Province:  m.Province,
			Regency:   m.Regency,
			District:  m.District,
			Village:   m.Village,
		})
	}
	return response
}

func (r *mosqueRepository) Show(id uint) (*models.MosqueResponse, error) {
	var mosque models.Mosque
	err := config.DB.Preload("Province").Preload("Regency").Preload("District").Preload("Village").Where("id = ?", id).First(&mosque).Error

	if err != nil {
		return nil, err
	}

	response := &models.MosqueResponse{
		ID:        mosque.ID,
		Name:      mosque.Name,
		Address:   mosque.Address,
		Photos:    mosque.Photos,
		CreatedAt: mosque.CreatedAt,
		UpdatedAt: mosque.UpdatedAt,
		Province:  mosque.Province,
		Regency:   mosque.Regency,
		District:  mosque.District,
		Village:   mosque.Village,
	}
	return response, nil
}
