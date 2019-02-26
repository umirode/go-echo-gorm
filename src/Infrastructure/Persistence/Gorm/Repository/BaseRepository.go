package Repository

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/go-rest/Database/Connection/Gorm"
)

type BaseRepository struct {
}

func (r *BaseRepository) GetGormDB() *gorm.DB {
	return Gorm.NewDatabase()
}
