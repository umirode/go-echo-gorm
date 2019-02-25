package Gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&Entity.Birthday{},
		&Entity.Notification{},
		&Entity.NotificationToken{},
		&Entity.RefreshToken{},
		&Entity.User{},
	)
}
