package database

import (
    "github.com/jinzhu/gorm"

    _ "github.com/jinzhu/gorm/dialects/mysql"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    _ "github.com/jinzhu/gorm/dialects/mssql"
    )

func NewDatabase(config *Config) (*gorm.DB, error) {
    dsn := NewDsnGenerator(config).GetDSN()

    return gorm.Open(config.Driver, dsn)
}
