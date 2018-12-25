package Controller

import (
	"github.com/labstack/echo"
	"github.com/umirode/go-rest/Http/Error"
	"github.com/umirode/go-rest/src/Application/Hydrator"
	"github.com/umirode/go-rest/src/Common"
	"github.com/umirode/go-rest/src/Domain/Service"
	"github.com/umirode/go-rest/src/Domain/Service/DTO"
	"net/http"
	"time"
)

type BirthdayController struct {
	BaseController

	BirthdayService  Service.IBirthdayService
	BirthdayHydrator Common.IHydrator
}

func NewBirthdayController(birthdayService Service.IBirthdayService, userService Service.IUserService) *BirthdayController {
	controller := &BirthdayController{
		BirthdayService:  birthdayService,
		BirthdayHydrator: &Hydrator.BirthdayHydrator{},
	}

	controller.UserService = userService

	return controller
}

func (c *BirthdayController) Index(context echo.Context) error {
	user, err := c.getCurrentUser(context)
	if err != nil {
		return err
	}

	birthdays, err := c.BirthdayService.GetAllForUser(user)
	if err != nil {
		return err
	}

	birthdaysMapArray := make([]map[string]interface{}, 0)
	for _, holiday := range birthdays {
		holidayMap, _ := c.BirthdayHydrator.Extract(holiday)
		birthdaysMapArray = append(birthdaysMapArray, holidayMap)
	}

	return context.JSON(http.StatusOK, birthdaysMapArray)
}

func (c *BirthdayController) Create(context echo.Context) error {
	data := new(struct {
		Name string `json:"name" validate:"required,max=100"`
		Date string `json:"date" validate:"required,date"`
	})

	if err := context.Bind(data); err != nil {
		return Error.NewRequestParsingError()
	}

	if err := context.Validate(data); err != nil {
		return err
	}

	user, err := c.getCurrentUser(context)
	if err != nil {
		return err
	}

	birthdayDTO := &DTO.BirthdayDTO{
		Name: data.Name,
	}

	birthdayDTO.Date, err = time.Parse(time.RFC822, data.Date)
	if err != nil {
		return err
	}

	err = c.BirthdayService.Create(birthdayDTO, user)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, nil)
}

func (c *BirthdayController) Update(context echo.Context) error {
	birthdayID, err := c.getParam(context, "birthday_id", "uint")
	if err != nil {
		return err
	}

	data := new(struct {
		Name string `json:"name" validate:"required,max=100"`
		Date string `json:"date" validate:"required,len=10,date"`
	})

	if err := context.Bind(data); err != nil {
		return Error.NewRequestParsingError()
	}

	if err := context.Validate(data); err != nil {
		return err
	}

	user, err := c.getCurrentUser(context)
	if err != nil {
		return err
	}

	birthday, err := c.BirthdayService.GetOneById(birthdayID.(uint))
	if err != nil {
		return err
	}

	birthdayDTO := &DTO.BirthdayDTO{
		Name: data.Name,
	}

	birthdayDTO.Date, err = time.Parse(time.RFC822, data.Date)
	if err != nil {
		return err
	}

	err = c.BirthdayService.Update(birthday, birthdayDTO, user)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, nil)
}

func (c *BirthdayController) Delete(context echo.Context) error {
	birthdayID, err := c.getParam(context, "birthday_id", "uint")
	if err != nil {
		return err
	}

	user, err := c.getCurrentUser(context)
	if err != nil {
		return err
	}

	birthday, err := c.BirthdayService.GetOneById(birthdayID.(uint))
	if err != nil {
		return err
	}

	err = c.BirthdayService.Delete(birthday, user)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, nil)
}
