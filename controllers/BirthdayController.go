package controllers

import (
	"github.com/labstack/echo"
	"github.com/umirode/go-rest/errors"
	"github.com/umirode/go-rest/models"
	"github.com/umirode/go-rest/services"
	"net/http"
)

type BirthdayController struct {
	BaseController

	BirthdayService services.IBirthdayService
}

func (c *BirthdayController) GetAllBirthdaysForUser(context echo.Context) error {
	claims, err := c.getTokenClaims(context)
	if err != nil {
		return err
	}
	userID := uint(claims["user_id"].(float64))

	birthdays, err := c.BirthdayService.GetAllBirthdaysForUser(userID)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, birthdays)
}

func (c *BirthdayController) CreateBirthdayByUser(context echo.Context) error {
	data := new(struct {
		Name string `json:"name" validate:"required,max=100"`
		Date string `json:"date" validate:"required,len=10,date"`
	})

	if err := context.Bind(data); err != nil {
		return errors.NewRequestParsingError()
	}

	if err := context.Validate(data); err != nil {
		return err
	}

	claims, err := c.getTokenClaims(context)
	if err != nil {
		return err
	}
	userID := uint(claims["user_id"].(float64))

	c.BirthdayService.CreateBirthdayByUser(userID, &models.BirthdayModel{
		Name: data.Name,
		Date: data.Date,
	})

	return nil
}

func (c *BirthdayController) UpdateBirthdayByUser(context echo.Context) error {
	birthdayID, err := c.getParam(context, "birthday_id", "uint")
	if err != nil {
		return err
	}

	data := new(struct {
		Name string `json:"name" validate:"required,max=100"`
		Date string `json:"date" validate:"required,len=10,date"`
	})

	if err := context.Bind(data); err != nil {
		return errors.NewRequestParsingError()
	}

	if err := context.Validate(data); err != nil {
		return err
	}

	claims, err := c.getTokenClaims(context)
	if err != nil {
		return err
	}
	userID := uint(claims["user_id"].(float64))

	return c.BirthdayService.UpdateBirthdayByUser(userID, birthdayID.(uint), &models.BirthdayModel{
		Name: data.Name,
		Date: data.Date,
	})
}

func (c *BirthdayController) DeleteBirthdayByUser(context echo.Context) error {
	birthdayID, err := c.getParam(context, "birthday_id", "uint")
	if err != nil {
		return err
	}

	claims, err := c.getTokenClaims(context)
	if err != nil {
		return err
	}
	userID := uint(claims["user_id"].(float64))

	return c.BirthdayService.DeleteBirthdayByUser(userID, birthdayID.(uint))
}
