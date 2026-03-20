package models

import "time"

type MosqueMemberRole string

const (
	MosqueAdmin     MosqueMemberRole = "admin"
	MosqueCommittee MosqueMemberRole = "committee"
	MosqueViewer    MosqueMemberRole = "viewer"
)

type MosqueMember struct {
	ID        uint             `json:"id" gorm:"primaryKey"`
	MosqueID  uint             `json:"mosque_id" gorm:"not null"`
	UserID    uint             `json:"user_id" gorm:"not null"`
	Role      MosqueMemberRole `json:"role" gorm:"type:varchar(20);default:'viewer'"`
	CreatedAt time.Time        `json:"created_at"`

	Mosque *Mosque `json:"mosque" gorm:"foreignKey:MosqueID"`
	User   *User   `json:"user" gorm:"foreignKey:UserID"`
}

func (MosqueMember) TableName() string {
	return "mosque_members"
}
