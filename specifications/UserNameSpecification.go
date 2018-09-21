package specifications

import (
	"github.com/jinzhu/gorm"
)

type UserNameSpecification struct {
	IDatabaseSpecification

	Name string
}

func (s *UserNameSpecification) GetForDatabase(db *gorm.DB) *gorm.DB {
	return db.Where("name = ?", s.Name)
}
