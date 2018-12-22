package Database

import (
	"github.com/Selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunMigrations(t *testing.T) {
	go_mocket.Catcher.Register()
	db, err := NewDatabase(&Config{
		Driver: go_mocket.DRIVER_NAME,
	})

	assert.Empty(t, err)

	assert.NotPanics(t, func() {
		RunMigrations(db)
	})
}
