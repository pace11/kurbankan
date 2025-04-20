package models

type UserCreateDTO struct {
	Email        string  `json:"email" binding:"required,email"`
	Password     string  `json:"password" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	Address      *string `json:"address"`
	Photos       *string `json:"photos"`
	ProvinceCode string  `json:"province_code" binding:"required"`
	RegencyCode  string  `json:"regency_code" binding:"required"`
	DistrictCode string  `json:"district_code" binding:"required"`
	VillageCode  string  `json:"village_code" binding:"required"`
}

type UserUpdateDTO struct {
	Email        string  `json:"email" binding:"required,email"`
	Password     string  `json:"password"`
	Name         string  `json:"name" binding:"required"`
	Address      *string `json:"address"`
	Photos       *string `json:"photos"`
	Role         *string `json:"role" binding:"required"`
	ProvinceCode string  `json:"province_code" binding:"required"`
	RegencyCode  string  `json:"regency_code" binding:"required"`
	DistrictCode string  `json:"district_code" binding:"required"`
	VillageCode  string  `json:"village_code" binding:"required"`
}
