package configs

import (
    "sync"
)

type DatabaseConfig struct {
    Driver   string
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
            Driver:   "sqlite3",
            Username: "root",
            Password: "",
            Host:     "127.0.0.1",
            Port:     3306,
            Database: "app.sqlite",
            Params:   "",
        }
    })

    return databaseConfig
}
