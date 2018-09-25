package controllers

import (
    "github.com/labstack/echo"
    "github.com/umirode/go-rest/models"
    "github.com/umirode/go-rest/response"
    "github.com/umirode/go-rest/services"
    )

type UserController struct {
    BaseController

    UserService services.IUserService
}

func (c *UserController) GetAllUsers(context echo.Context) error {
    users := c.UserService.GetAllUsers()

    return response.SendResponseJson(context, "success", "", users)
}

func (c *UserController) GetSingleUser(context echo.Context) error {
    id := c.getParam(context, "id", "uint").(uint)
    if id == 0 {
        return c.getQueryParsingError()
    }

    user := c.UserService.GetUserByID(id)
    if user == nil{
       return c.getNotExistsError()
    }

    return response.SendResponseJson(context, "success", "", user)
}

func (c *UserController) CreateUser(context echo.Context) error {
    user := new(models.UserModel)

    err := context.Bind(user)
    if err != nil {
        return err
    }

    err = c.UserService.CreateUser(user)
    if err != nil {
        return err
    }

    return response.SendResponseJson(context, "success", "", nil)
}

func (c *UserController) UpdateUser(context echo.Context) error {
    id := c.getParam(context, "id", "uint").(uint)
    if id == 0 {
        return c.getQueryParsingError()
    }

    user := new(models.UserModel)

    err := context.Bind(user)
    if err != nil {
        return err
    }

    err = c.UserService.UpdateUser(id, user)
    if err != nil {
        return err
    }

    return response.SendResponseJson(context, "success", "", nil)
}

func (c *UserController) DeleteUser(context echo.Context) error {
    id := c.getParam(context, "id", "uint").(uint)
    if id == 0 {
        return c.getQueryParsingError()
    }

    err := c.UserService.DeleteUser(id)
    if err != nil {
        return err
    }

    return response.SendResponseJson(context, "success", "", nil)
}
