package controllers

import (
	"github.com/labstack/echo"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) GetAllUsers(c echo.Context) error {

}

func (uc *UserController) GetSingleUser(c echo.Context) error {

}

func (uc *UserController) CreateUser(c echo.Context) error {

}

func (uc *UserController) UpdateUser(c echo.Context) error {

}

func (uc *UserController) DeleteUser(c echo.Context) error {

}
