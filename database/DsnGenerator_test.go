package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDsnGenerator(t *testing.T) {
	assert.NotEmpty(t, NewDsnGenerator(&Config{}))
}

func TestDsnGenerator_GetDSN(t *testing.T) {
	drivers := [...]string{
		"mysql",
		"postgres",
		"sqlite3",
		"mssql",
	}

	config := &Config{
		Database: "test",
	}

	for _, driver := range drivers {
		config.Driver = driver

		dsnGenerator := NewDsnGenerator(config)

		assert.NotEmpty(t, dsnGenerator.GetDSN())
	}
}
