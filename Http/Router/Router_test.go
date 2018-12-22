package Router

import (
	"github.com/Selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"github.com/umirode/go-rest/Database"
	"testing"
)

func TestNewRouter(t *testing.T) {
	go_mocket.Catcher.Register()
	db, _ := Database.NewDatabase(&Database.Config{
		Driver: go_mocket.DRIVER_NAME,
	})

	router := NewRouter(db, true)

	assert.NotEmpty(t, router)
}
