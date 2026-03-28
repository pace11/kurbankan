package models

type UserCreatePayload struct {
	Email        string  `json:"email" binding:"required,email"`
	Password     string  `json:"password" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	Gender       *Gender `json:"gender"`
	Address      *string `json:"address"`
	Photos       *string `json:"photos"`
	ProvinceCode string  `json:"province_code" binding:"required"`
	RegencyCode  string  `json:"regency_code" binding:"required"`
	DistrictCode string  `json:"district_code" binding:"required"`
	VillageCode  string  `json:"village_code" binding:"required"`
}

type UserUpdatePayload struct {
	Email        string  `json:"email" binding:"required,email"`
	Password     string  `json:"password"`
	Name         string  `json:"name" binding:"required"`
	Gender       *Gender `json:"gender"`
	Address      *string `json:"address"`
	Photos       *string `json:"photos"`
	Role         *string `json:"role" binding:"required"`
	ProvinceCode string  `json:"province_code" binding:"required"`
	RegencyCode  string  `json:"regency_code" binding:"required"`
	DistrictCode string  `json:"district_code" binding:"required"`
	VillageCode  string  `json:"village_code" binding:"required"`
}

type LoginPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type TransactionPayload struct {
	MosqueID         uint    `json:"mosque_id" binding:"required"`
	QurbanOfferingID uint    `json:"qurban_offering_id" binding:"required"`
	ParticipantID    uint    `json:"participant_id" binding:"required"`
	Amount           float64 `json:"amount" binding:"required"`
}
