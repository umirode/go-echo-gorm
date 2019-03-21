package Middleware

import (
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestNewLoggerMiddleware(t *testing.T) {
	middleware := NewLoggerMiddleware()

	assert.NotNil(t, middleware)
}

func TestLoggerMiddleware_Middleware(t *testing.T) {
	middleware := NewLoggerMiddleware()
	middlewareFunc := middleware.Middleware(func(context echo.Context) error {
		return nil
	})

	assert.NotNil(t, middleware)
	assert.NotNil(t, middlewareFunc)
}
