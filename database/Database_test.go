package database

import (
	"github.com/Selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	go_mocket.Catcher.Register()
	db, err := NewDatabase(&Config{
		Driver: go_mocket.DRIVER_NAME,
	})

	assert.NotEmpty(t, db)
	assert.NoError(t, err)
}

func TestNewDatabase_Error(t *testing.T) {
	db, err := NewDatabase(&Config{
		Driver: "test",
	})

	assert.Empty(t, db)
	assert.Error(t, err)
}
