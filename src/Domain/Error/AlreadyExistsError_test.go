package Error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAlreadyExistsError(t *testing.T) {
	err := NewAlreadyExistsError()

	assert.NotNil(t, err)
}

func TestAlreadyExistsError_Status(t *testing.T) {
	err := NewAlreadyExistsError()

	assert.NotNil(t, err)
	assert.Equal(t, 409, err.Status())
}

func TestAlreadyExistsError_Error(t *testing.T) {
	err := NewAlreadyExistsError()

	assert.NotNil(t, err)
	assert.Equal(t, "Already exists", err.Error())
}
