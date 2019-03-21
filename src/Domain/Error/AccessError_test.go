package Error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccessError(t *testing.T) {
	err := NewAccessError()

	assert.NotNil(t, err)
}

func TestAccessError_Status(t *testing.T) {
	err := NewAccessError()

	assert.NotNil(t, err)
	assert.Equal(t, 403, err.Status())
}

func TestAccessError_Error(t *testing.T) {
	err := NewAccessError()

	assert.NotNil(t, err)
	assert.Equal(t, "Access error", err.Error())
}
