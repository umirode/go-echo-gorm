package database

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func NewDatabase(config *Config) (*gorm.DB, error) {
	dsn := NewDsnGenerator(config).GetDSN()

	db, err := gorm.Open(config.Driver, dsn)

	db.LogMode(config.Debug)

	return db, err
}
