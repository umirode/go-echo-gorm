package Error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRequestParsingError(t *testing.T) {
	err := NewRequestParsingError()

	assert.NotNil(t, err)
}

func TestRequestParsingError_Status(t *testing.T) {
	err := NewRequestParsingError()

	assert.NotNil(t, err)
	assert.Equal(t, 400, err.Status())
}

func TestRequestParsingError_Error(t *testing.T) {
	err := NewRequestParsingError()

	assert.NotNil(t, err)
	assert.Equal(t, "Parsing request error", err.Error())
}
