package Response

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResponse(t *testing.T) {
	response := NewResponse(200, "test", "test")

	assert.NotNil(t, response)
	assert.Equal(t, 200, response.Status)
	assert.Equal(t, "test", response.Data)
	assert.Equal(t, "test", response.Message)
}
