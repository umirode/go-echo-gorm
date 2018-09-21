package configs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetServerConfig(t *testing.T) {
	assert.NotEmpty(t, GetServerConfig())
}
