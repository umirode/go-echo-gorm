package Entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRefreshToken(t *testing.T) {
	token := &RefreshToken{
		ID:        1,
		Token:     "test",
		ExpiresAt: 10,
		OwnerID:   1,
	}

	assert.NotNil(t, token)
	assert.Equal(t, uint(1), token.ID)
	assert.Equal(t, "test", token.Token)
	assert.Equal(t, int64(10), token.ExpiresAt)
	assert.Equal(t, uint(1), token.OwnerID)
}
