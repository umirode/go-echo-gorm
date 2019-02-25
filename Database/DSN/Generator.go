package DSN

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/umirode/go-rest/Database"
)

type Generator struct {
	Config *Database.Config
}

func NewGenerator(config *Database.Config) *Generator {
	return &Generator{
		Config: config,
	}
}

func (d *Generator) GetDSN() string {
	switch d.Config.Driver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&%s", d.Config.Username, d.Config.Password, d.Config.Host, d.Config.Port, d.Config.Database, d.Config.Params)
	case "postgres":
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s %s", d.Config.Host, d.Config.Port, d.Config.Username, d.Config.Database, d.Config.Password, d.Config.Params)
	case "sqlite3":
		return fmt.Sprintf("%s", d.Config.Database)
	case "mssql":
		return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&%s", d.Config.Username, d.Config.Password, d.Config.Host, d.Config.Port, d.Config.Database, d.Config.Params)
	default:
		logrus.Fatalf("Driver does not support: %s", d.Config.Driver)
	}

	return ""
}
