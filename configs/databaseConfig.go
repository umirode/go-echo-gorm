package configs

import (
	"os"
	"strconv"
	"sync"
)

type DatabaseConfig struct {
	Driver   string
	Debug    bool
	Username string
	Password string
	Host     string
	Port     uint
	Database string
	Params   string
}

var databaseConfigOnce sync.Once
var databaseConfig *DatabaseConfig

func GetDatabaseConfig() *DatabaseConfig {
	databaseConfigOnce.Do(func() {
		port, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))
		debug, _ := strconv.ParseBool(os.Getenv("DATABASE_DEBUG"))

		databaseConfig = &DatabaseConfig{
			Driver:   os.Getenv("DATABASE_DRIVER"),
			Debug:    debug,
			Username: os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     uint(port),
			Database: os.Getenv("DATABASE_NAME"),
			Params:   os.Getenv("DATABASE_PARAMS"),
		}
	})

	return databaseConfig
}
