package Config

import (
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
		databaseConfig = &DatabaseConfig{
			Driver:   GetEnv("DATABASE_DRIVER", "string").(string),
			Debug:    GetEnv("DATABASE_DEBUG", "bool").(bool),
			Username: GetEnv("DATABASE_USER", "string").(string),
			Password: GetEnv("DATABASE_PASSWORD", "string").(string),
			Host:     GetEnv("DATABASE_HOST", "string").(string),
			Port:     GetEnv("DATABASE_PORT", "uint").(uint),
			Database: GetEnv("DATABASE_NAME", "string").(string),
			Params:   GetEnv("DATABASE_PARAMS", "string").(string),
		}
	})

	return databaseConfig
}
