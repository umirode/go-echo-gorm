package ValueObject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNotification(t *testing.T) {
	notification := NewNotification("test", "test")

	assert.NotNil(t, notification)
	assert.Equal(t, "test", notification.Title)
	assert.Equal(t, "test", notification.Message)
}
