package Gorm

import (
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/umirode/go-rest/Database/Connection"
	"github.com/umirode/go-rest/Database/DSN"
)

var databaseOnce sync.Once
var db *gorm.DB

func NewDatabase() *gorm.DB {
	databaseOnce.Do(func() {
		config := Connection.GetConfig("gorm")

		dsn := DSN.NewGenerator(config).GetDSN()

		err := *new(error)
		db, err = gorm.Open(config.Driver, dsn)
		if err != nil {
			logrus.Fatal(err)
		}

		db.DB().SetConnMaxLifetime(time.Minute * 5)
		db.DB().SetMaxIdleConns(0)
		db.DB().SetMaxOpenConns(5)

		db.LogMode(config.Debug)

		runMigrations(db)
	})

	return db
}
