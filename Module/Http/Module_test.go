package Http

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	Module2 "github.com/umirode/go-rest/Module"
)

func TestNewModule(t *testing.T) {
	module := NewModule()

	assert.NotNil(t, module)
	assert.Implements(t, (*Module2.IModule)(nil), module)
}

func TestModule_Init(t *testing.T) {
	module := NewModule()

	assert.NotNil(t, module)
	assert.NotPanics(t, func() {
		module.Init(&sync.WaitGroup{})
	})
}

func TestModule_Close(t *testing.T) {
	module := NewModule()

	assert.NotNil(t, module)
	assert.NotPanics(t, func() {
		module.Close(&sync.WaitGroup{})
	})
}
