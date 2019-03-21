package ValueObject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJWT(t *testing.T) {
	token := NewJWT("test", 10)

	assert.NotNil(t, token)
	assert.Equal(t, "test", token.Token)
	assert.Equal(t, int64(10), token.ExpiresAt)
}
