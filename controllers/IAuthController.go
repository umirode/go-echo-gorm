package controllers

import (
	"github.com/labstack/echo"
)

type IAuthController interface {
	Login(context echo.Context) error
	Signup(context echo.Context) error
	RefreshToken(context echo.Context) error
	Logout(context echo.Context) error
	ResetPassword(context echo.Context) error
}
