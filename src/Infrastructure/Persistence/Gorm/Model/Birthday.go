package Model

import (
	"time"
)

type Birthday struct {
	ID uint `gorm:"primary_key" json:"id"`

	UserID uint      `gorm:"not null;" json:"user_id"`
	Name   string    `gorm:"not null;size:100" json:"name"`
	Date   time.Time `gorm:"not null" json:"date"`
}
