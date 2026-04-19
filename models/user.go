package models

import (
	"time"
)

type PlatformRole string

const (
	PlatformRoleOwner   PlatformRole = "owner"
	PlatformRoleAdmin   PlatformRole = "admin"
	PlatformRoleSupport PlatformRole = "support"
)

type UserResponse struct {
	ID           uint          `json:"id"`
	Email        string        `json:"email"`
	PlatformRole *PlatformRole `json:"platform_role"`
	CreatedAt    *time.Time    `json:"created_at"`
	UpdatedAt    *time.Time    `json:"updated_at"`
}

type User struct {
	ID           uint          `json:"id" gorm:"primaryKey"`
	Email        string        `json:"email" gorm:"unique;not null" binding:"required"`
	Password     string        `json:"password" gorm:"type:text;not null"`
	PlatformRole *PlatformRole `json:"platform_role" gorm:"type:varchar(20)"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func ToUserResponse(u *User) *UserResponse {
	if u == nil {
		return nil
	}
	return &UserResponse{
		ID:           u.ID,
		Email:        u.Email,
		PlatformRole: u.PlatformRole,
		CreatedAt:    &u.CreatedAt,
		UpdatedAt:    &u.UpdatedAt,
	}
}

func IsPlatformRoleOwner(role string) bool {
	return role == string(PlatformRoleOwner)
}

func IsPlatformRoleAdmin(role string) bool {
	return role == string(PlatformRoleAdmin)
}

func IsPlatformRoleSupport(role string) bool {
	return role == string(PlatformRoleSupport)
}
