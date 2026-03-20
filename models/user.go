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
