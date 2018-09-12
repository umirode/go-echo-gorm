package database

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var dbConnection *gorm.DB

func OpenDBConnection(
	driver string,
	database string,
	host string,
	port uint,
	username string,
	password string,
	params string,
) {
	dsn := getDSN(
		driver,
		database,
		host,
		port,
		username,
		password,
		params,
	)

	var err error
	dbConnection, err = gorm.Open(driver, dsn)

	if err != nil {
		logrus.Errorf("Database connection open error: %s", err)
	}
}

func CloseDBConnection() {
	if dbConnection != nil {
		dbConnection.Close()
	}
}

func GetDBConnection() *gorm.DB {
	if dbConnection != nil {
		return dbConnection
	}

	logrus.Error("DB connection not opened")

	return nil
}

func getDSN(
	driver string,
	database string,
	host string,
	port uint,
	username string,
	password string,
	params string,
) string {
	switch driver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", username, password, host, port, database, params)
		break
	case "postgres":
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s %s", host, port, username, database, password, params)
		break
	case "sqlite3":
		return fmt.Sprintf("%s", database)
		break
	case "mssql":
		return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&%s", username, password, host, port, database, params)
		break
	}

	return ""
}
