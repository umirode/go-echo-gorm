package Error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBirthdaysLimitError(t *testing.T) {
	err := NewBirthdaysLimitError()

	assert.NotNil(t, err)
}

func TestBirthdaysLimitError_Status(t *testing.T) {
	err := NewBirthdaysLimitError()

	assert.NotNil(t, err)
	assert.Equal(t, 400, err.Status())
}

func TestBirthdaysLimitError_Error(t *testing.T) {
	err := NewBirthdaysLimitError()

	assert.NotNil(t, err)
	assert.Equal(t, "Type convert error", err.Error())
}
