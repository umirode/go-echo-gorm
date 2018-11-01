package controllers

import (
	"github.com/labstack/echo"
)

type IBirthdayController interface {
	GetAllBirthdaysForUser(context echo.Context) error
	CreateBirthdayByUser(context echo.Context) error
	UpdateBirthdayByUser(context echo.Context) error
	DeleteBirthdayByUser(context echo.Context) error
}
