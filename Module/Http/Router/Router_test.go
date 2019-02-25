package Router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	router := NewRouter(true)

	assert.NotEmpty(t, router)
}
