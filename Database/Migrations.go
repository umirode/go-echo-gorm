package Database

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/go-rest/src/Infrastructure/Persistence/Gorm/Model"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&Model.User{},
		&Model.Birthday{},
	)
}
