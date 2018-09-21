package database

import (
    "testing"
    "github.com/Selvatico/go-mocket"
    "github.com/stretchr/testify/assert"
)

func TestNewDatabase(t *testing.T) {
    go_mocket.Catcher.Register()
    db, err := NewDatabase(&Config{
        Driver: go_mocket.DRIVER_NAME,
    })

    assert.NotEmpty(t, db)
    assert.Empty(t, err)
}
