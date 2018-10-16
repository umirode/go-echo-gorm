package router

import (
	"github.com/Selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"github.com/umirode/go-rest/database"
	"testing"
)

func TestNewRouter(t *testing.T) {
	go_mocket.Catcher.Register()
	db, _ := database.NewDatabase(&database.Config{
		Driver: go_mocket.DRIVER_NAME,
	})

	router := NewRouter(db, true)

	assert.NotEmpty(t, router)
}
