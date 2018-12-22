package Middleware

import (
	"github.com/labstack/echo"
)

type IMiddleware interface {
	Middleware(next echo.HandlerFunc) echo.HandlerFunc
}
