package Entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNotificationToken(t *testing.T) {
	token := &NotificationToken{
		ID:      1,
		Token:   "test",
		OwnerID: 1,
	}

	assert.NotNil(t, token)
	assert.Equal(t, uint(1), token.ID)
	assert.Equal(t, "test", token.Token)
	assert.Equal(t, uint(1), token.OwnerID)
}
