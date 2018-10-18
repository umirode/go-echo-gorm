package middlewares

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewJWTMiddleware(t *testing.T) {
	middleware := NewJWTMiddleware("test")

	assert.NotEmpty(t, middleware)
}

func TestJWTMiddleware_Middleware(t *testing.T) {
	middleware := NewJWTMiddleware("test")
	middlewareFunc := middleware.Middleware(func(context echo.Context) error {
		return nil
	})

	assert.NotEmpty(t, middleware)
	assert.NotEmpty(t, middlewareFunc)
}
