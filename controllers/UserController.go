package controllers

import (
	"github.com/labstack/echo"
	"github.com/umirode/go-rest/errors"
	"github.com/umirode/go-rest/models"
	"github.com/umirode/go-rest/services"
	"net/http"
)

type UserController struct {
	BaseController

	UserService services.IUserService
}

func (c *UserController) GetAllUsers(context echo.Context) error {
	users, err := c.UserService.GetAllUsers()
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, users)
}

func (c *UserController) GetSingleUser(context echo.Context) error {
	id, err := c.getParam(context, "id", "uint")
	if err != nil {
		return err
	}

	user, err := c.UserService.GetUserByID(id.(uint))
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, user)
}

func (c *UserController) CreateUser(context echo.Context) error {
	user := new(models.UserModel)

	err := context.Bind(user)
	if err != nil {
		return errors.NewRequestParsingError()
	}

	err = c.UserService.CreateUser(user)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, nil)
}

func (c *UserController) UpdateUser(context echo.Context) error {
	id, err := c.getParam(context, "id", "uint")
	if err != nil {
		return err
	}

	user := new(models.UserModel)

	err = context.Bind(user)
	if err != nil {
		return errors.NewRequestParsingError()
	}

	err = c.UserService.UpdateUser(id.(uint), user)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, nil)
}

func (c *UserController) DeleteUser(context echo.Context) error {
	id, err := c.getParam(context, "id", "uint")
	if err != nil {
		return err
	}

	err = c.UserService.DeleteUser(id.(uint))
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, nil)
}
