package configs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDatabaseConfig(t *testing.T) {
	assert.NotEmpty(t, GetDatabaseConfig())
}
