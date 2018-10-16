package database

import (
	"github.com/jinzhu/gorm"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func NewDatabase(config *Config) (*gorm.DB, error) {
	dsn := NewDsnGenerator(config).GetDSN()

	db, err := gorm.Open(config.Driver, dsn)
	if err != nil {
		return nil, err
	}

	db.DB().SetConnMaxLifetime(time.Minute * 5)
	db.DB().SetMaxIdleConns(0)
	db.DB().SetMaxOpenConns(5)

	db.LogMode(config.Debug)

	return db, nil
}
