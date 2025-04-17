package models

import (
	"time"
)

type UserRole string

const (
	Admin        UserRole = "admin"
	MosqueMember UserRole = "mosque_member"
	UserMember   UserRole = "user_member"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"type:text;not null"`
	Role      UserRole  `json:"role" gorm:"type:enum('admin','mosque_member','user_member');default:'user_member';not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
