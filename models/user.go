package models

import (
	"time"
)

type UserRole string

const (
	Admin            UserRole = "admin"
	RoleMosqueMember UserRole = "mosque_member"
	RoleUserMember   UserRole = "user_member"
)

type UserResponse struct {
	ID        uint       `json:"id"`
	Email     string     `json:"email"`
	Role      UserRole   `json:"role"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique;not null" binding:"required"`
	Password  string    `json:"password" gorm:"type:text;not null"`
	Role      UserRole  `json:"role" gorm:"type:varchar(50);default:'user_member';not null" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func ToUserResponse(u *User) *UserResponse {
	if u == nil {
		return nil
	}
	return &UserResponse{
		ID:    u.ID,
		Email: u.Email,
		Role:  u.Role,
	}
}
