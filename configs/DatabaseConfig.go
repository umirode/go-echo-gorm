package configs

import (
	"github.com/umirode/go-rest/Env"
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
			Driver:   Env.GetEnv("DATABASE_DRIVER", "string").(string),
			Debug:    Env.GetEnv("DATABASE_DEBUG", "bool").(bool),
			Username: Env.GetEnv("DATABASE_USER", "string").(string),
			Password: Env.GetEnv("DATABASE_PASSWORD", "string").(string),
			Host:     Env.GetEnv("DATABASE_HOST", "string").(string),
			Port:     Env.GetEnv("DATABASE_PORT", "uint").(uint),
			Database: Env.GetEnv("DATABASE_NAME", "string").(string),
			Params:   Env.GetEnv("DATABASE_PARAMS", "string").(string),
		}
	})

	return databaseConfig
}
