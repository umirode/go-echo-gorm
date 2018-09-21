package controllers

import (
	"github.com/labstack/echo"
)

type IUserController interface {
	GetAllUsers(context echo.Context) error
	GetSingleUser(context echo.Context) error
	CreateUser(context echo.Context) error
	UpdateUser(context echo.Context) error
	DeleteUser(context echo.Context) error
}
