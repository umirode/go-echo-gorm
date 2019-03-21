package DSN

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/umirode/go-rest/Database"
)

func TestNewGenerator(t *testing.T) {
	assert.NotNil(t, NewGenerator(&Database.Config{}))
}

func TestDsnGenerator_GetDSN(t *testing.T) {
	drivers := [...]string{
		"mysql",
		"postgres",
		"sqlite3",
		"mssql",
	}

	config := &Database.Config{
		Database: "test",
	}

	for _, driver := range drivers {
		config.Driver = driver

		dsnGenerator := NewGenerator(config)

		assert.NotNil(t, dsnGenerator)
		assert.NotEmpty(t, dsnGenerator.GetDSN())
	}
}
