package database

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/go-rest/models"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&models.UserModel{},
	)
}
