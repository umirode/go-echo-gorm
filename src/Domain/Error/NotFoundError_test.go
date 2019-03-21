package Error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError()

	assert.NotNil(t, err)
}

func TestNotFoundError_Status(t *testing.T) {
	err := NewNotFoundError()

	assert.NotNil(t, err)
	assert.Equal(t, 404, err.Status())
}

func TestNotFoundError_Error(t *testing.T) {
	err := NewNotFoundError()

	assert.NotNil(t, err)
	assert.Equal(t, "Not found", err.Error())
}
