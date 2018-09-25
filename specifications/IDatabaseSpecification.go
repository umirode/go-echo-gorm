package specifications

import (
    "github.com/jinzhu/gorm"
)

type IDatabaseSpecification interface {
    GetForDatabase(db *gorm.DB) *gorm.DB
}
