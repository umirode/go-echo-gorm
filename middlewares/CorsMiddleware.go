package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type CorsMiddleware struct{}

func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{}
}

func (m *CorsMiddleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.CORS()(next)
}
