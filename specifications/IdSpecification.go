package specifications

import (
	"github.com/jinzhu/gorm"
)

type IdSpecification struct {
	IDatabaseSpecification

	Id uint
}

func (s *IdSpecification) GetForDatabase(db *gorm.DB) *gorm.DB {
	return db.Where("id = ?", s.Id)
}
