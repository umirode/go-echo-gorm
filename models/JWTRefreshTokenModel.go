package models

import (
	"time"
)

type JWTRefreshTokenModel struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UserID uint   `gorm:"not null;" json:"user_id"`
	Token  string `gorm:"not null;size:255" json:"token"`
	UserIP string `gorm:"not null;size:40" json:"user_ip"`
}

func (m *JWTRefreshTokenModel) TableName() string {
	return "jwt_refresh_tokens"
}
