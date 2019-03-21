package Helper

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvBool(t *testing.T) {
	err := os.Setenv("TEST", "true")
	assert.Nil(t, err)

	value := GetEnv("TEST", "bool").(bool)

	assert.Equal(t, true, value)
}

func TestGetEnvInt(t *testing.T) {
	err := os.Setenv("TEST", "10")
	assert.Nil(t, err)

	value := GetEnv("TEST", "int").(int)
	assert.Equal(t, 10, value)
}

func TestGetEnvInt64(t *testing.T) {
	err := os.Setenv("TEST", "10")
	assert.Nil(t, err)

	value := GetEnv("TEST", "int64").(int64)
	assert.Equal(t, int64(10), value)
}

func TestGetEnvUint64(t *testing.T) {
	err := os.Setenv("TEST", "10")
	assert.Nil(t, err)

	value := GetEnv("TEST", "uint64").(uint64)
	assert.Equal(t, uint64(10), value)
}

func TestGetEnvUint(t *testing.T) {
	err := os.Setenv("TEST", "10")
	assert.Nil(t, err)

	value := GetEnv("TEST", "uint").(uint)
	assert.Equal(t, uint(10), value)
}

func TestGetEnvString(t *testing.T) {
	err := os.Setenv("TEST", "test")
	assert.Nil(t, err)

	value := GetEnv("TEST", "string").(string)
	assert.Equal(t, "test", value)
}
