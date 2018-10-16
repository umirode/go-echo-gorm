package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type LoggerMiddleware struct {
	IMiddleware
}

func NewLoggerMiddleware() *LoggerMiddleware {
	return &LoggerMiddleware{}
}

func (m *LoggerMiddleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.Logger()(next)
}
