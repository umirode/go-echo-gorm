package models

import (
	"time"
)

type BirthdayModel struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UserID uint   `gorm:"not null;" json:"user_id"`
	Name   string `gorm:"not null;size:100" json:"name"`
	Date   string `gorm:"not null;size:10" json:"date"`
}

func (m *BirthdayModel) TableName() string {
	return "birthdays"
}
