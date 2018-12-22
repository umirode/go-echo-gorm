package Middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/umirode/go-rest/Http/Error"
)

type JWTMiddleware struct {
	Secret string
}

func NewJWTMiddleware(secret string) *JWTMiddleware {
	return &JWTMiddleware{
		Secret: secret,
	}
}

func (m *JWTMiddleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(m.Secret),
		ErrorHandler: func(e error) error {
			return Error.NewInvalidTokenError()
		},
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  "Bearer",
	})(next)
}
