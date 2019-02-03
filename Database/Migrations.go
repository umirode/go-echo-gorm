package Database

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&Entity.User{},
		&Entity.Birthday{},
		&Entity.RefreshToken{},
	)
}
