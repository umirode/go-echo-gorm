package Error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInvalidError(t *testing.T) {
	err := NewInvalidError()

	assert.NotNil(t, err)
}

func TestInvalidError_Status(t *testing.T) {
	err := NewInvalidError()

	assert.NotNil(t, err)
	assert.Equal(t, 400, err.Status())
}

func TestInvalidError_Error(t *testing.T) {
	err := NewInvalidError()

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid", err.Error())
}
