package Config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetServerConfig(t *testing.T) {
	err := os.Setenv("SERVER_HOST", "test")
	assert.Nil(t, err)
	err = os.Setenv("SERVER_PORT", "10")
	assert.Nil(t, err)
	err = os.Setenv("SERVER_DEBUG", "true")
	assert.Nil(t, err)

	config := GetServerConfig()

	assert.NotNil(t, config)
	assert.Equal(t, "test", config.Host)
	assert.Equal(t, uint(10), config.Port)
	assert.Equal(t, true, config.Debug)
}
