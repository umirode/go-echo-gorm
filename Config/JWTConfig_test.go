package Config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetJWTConfig(t *testing.T) {
	err := os.Setenv("JWT_ACCESS_SECRET", "test")
	assert.Nil(t, err)
	err = os.Setenv("JWT_ACCESS_LIFE_TIME", "10")
	assert.Nil(t, err)
	err = os.Setenv("JWT_REFRESH_SECRET", "test")
	assert.Nil(t, err)
	err = os.Setenv("JWT_REFRESH_LIFE_TIME", "10")
	assert.Nil(t, err)

	config := GetJWTConfig()

	assert.NotNil(t, config)
	assert.Equal(t, "test", config.AccessTokenSecret)
	assert.Equal(t, int64(10), config.AccessTokenLifeTime)
	assert.Equal(t, "test", config.RefreshTokenSecret)
	assert.Equal(t, int64(10), config.RefreshTokenLifeTime)
}
