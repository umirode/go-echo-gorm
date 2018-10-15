package middlewares

import (
	"github.com/labstack/echo"
)

type IMiddleware interface {
	Middleware(next echo.HandlerFunc) echo.HandlerFunc
}
