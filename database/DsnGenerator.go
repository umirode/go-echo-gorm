package database

import (
    "fmt"
)

type DsnGenerator struct {
    Config *Config
}

func NewDsnGenerator(config *Config) *DsnGenerator {
    return &DsnGenerator{
        Config: config,
    }
}

func (d *DsnGenerator) GetDSN() string {
    switch d.Config.Driver {
    case "mysql":
        return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", d.Config.Username, d.Config.Password, d.Config.Host, d.Config.Port, d.Config.Database, d.Config.Params)
        break
    case "postgres":
        return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s %s", d.Config.Host, d.Config.Port, d.Config.Username, d.Config.Database, d.Config.Password, d.Config.Params)
        break
    case "sqlite3":
        return fmt.Sprintf("%s", d.Config.Database)
        break
    case "mssql":
        return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&%s", d.Config.Username, d.Config.Password, d.Config.Host, d.Config.Port, d.Config.Database, d.Config.Params)
        break
    }

    return ""
}
