package Entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user := &User{
		ID:           1,
		Email:        "test",
		PasswordHash: "test",
	}

	assert.NotNil(t, user)
	assert.Equal(t, uint(1), user.ID)
	assert.Equal(t, "test", user.Email)
	assert.Equal(t, "test", user.PasswordHash)
}
