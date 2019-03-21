package Middleware

import (
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestNewCorsMiddleware(t *testing.T) {
	middleware := NewCorsMiddleware()

	assert.NotNil(t, middleware)
}

func TestCorsMiddleware_Middleware(t *testing.T) {
	middleware := NewCorsMiddleware()
	middlewareFunc := middleware.Middleware(func(context echo.Context) error {
		return nil
	})

	assert.NotNil(t, middleware)
	assert.NotNil(t, middlewareFunc)
}
