package models

import (
    "time"
)

type UserModel struct {
    ID        uint       `gorm:"primary_key" json:"id"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
    DeletedAt *time.Time `sql:"index" json:"deleted_at"`

    Name string `gorm:"size:255" json:"name"`
}