package Error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInvalidTokenError(t *testing.T) {
	err := NewInvalidTokenError()

	assert.NotNil(t, err)
}

func TestInvalidTokenError_Status(t *testing.T) {
	err := NewInvalidTokenError()

	assert.NotNil(t, err)
	assert.Equal(t, 401, err.Status())
}

func TestInvalidTokenError_Error(t *testing.T) {
	err := NewInvalidTokenError()

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid or expired jwt", err.Error())
}
