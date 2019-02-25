package v1

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/umirode/go-rest/Module/Http/Controller"
	"github.com/umirode/go-rest/Module/Http/Error"
	"github.com/umirode/go-rest/src/Application/Hydrator"
	"github.com/umirode/go-rest/src/Common"
	"github.com/umirode/go-rest/src/Domain/Service"
	"github.com/umirode/go-rest/src/Domain/Service/DTO"
)

type BirthdayController struct {
	Controller.BaseController

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

func (c *BirthdayController) GetAll(context echo.Context) error {
	user, err := c.GetCurrentUser(context)
	if err != nil {
		return err
	}

	birthdays, err := c.BirthdayService.GetAllByUser(user)
	if err != nil {
		return err
	}

	birthdaysMapArray := make([]map[string]interface{}, 0)
	for _, birthday := range birthdays {
		birthdayMap, _ := c.BirthdayHydrator.Extract(birthday)
		birthdaysMapArray = append(birthdaysMapArray, birthdayMap)
	}

	return c.Response(context, http.StatusOK, birthdaysMapArray, "")
}

func (c *BirthdayController) GetOne(context echo.Context) error {
	birthdayID, err := c.GetParam(context, "birthday_id", "uint")
	if err != nil {
		return err
	}

	user, err := c.GetCurrentUser(context)
	if err != nil {
		return err
	}

	birthday, err := c.BirthdayService.GetOneByIdAndUser(birthdayID.(uint), user)
	if err != nil {
		return err
	}

	birthdayMap, _ := c.BirthdayHydrator.Extract(birthday)

	return c.Response(context, http.StatusOK, birthdayMap, "")
}

func (c *BirthdayController) Create(context echo.Context) error {
	data := new(struct {
		Name  string `json:"name" validate:"required,max=20"`
		Month uint   `json:"month" validate:"required,max=12"`
		Day   uint   `json:"day" validate:"required,max=31"`
	})

	if err := context.Bind(data); err != nil {
		return Error.NewRequestParsingError()
	}

	if err := context.Validate(data); err != nil {
		return err
	}

	user, err := c.GetCurrentUser(context)
	if err != nil {
		return err
	}

	birthdayDTO := &DTO.BirthdayDTO{
		Name:  data.Name,
		Month: data.Month,
		Day:   data.Day,
	}

	err = c.BirthdayService.Create(birthdayDTO, user)
	if err != nil {
		return err
	}

	return c.Response(context, http.StatusCreated, nil, "")
}

func (c *BirthdayController) Update(context echo.Context) error {
	birthdayID, err := c.GetParam(context, "birthday_id", "uint")
	if err != nil {
		return err
	}

	data := new(struct {
		Name  string `json:"name" validate:"required,max=20"`
		Month uint   `json:"month" validate:"required,max=12"`
		Day   uint   `json:"day" validate:"required,max=31"`
	})

	if err := context.Bind(data); err != nil {
		return Error.NewRequestParsingError()
	}

	if err := context.Validate(data); err != nil {
		return err
	}

	user, err := c.GetCurrentUser(context)
	if err != nil {
		return err
	}

	birthday, err := c.BirthdayService.GetOneByIdAndUser(birthdayID.(uint), user)
	if err != nil {
		return err
	}

	birthdayDTO := &DTO.BirthdayDTO{
		Name:  data.Name,
		Month: data.Month,
		Day:   data.Day,
	}

	err = c.BirthdayService.Update(birthday, birthdayDTO, user)
	if err != nil {
		return err
	}

	return c.Response(context, http.StatusOK, nil, "")
}

func (c *BirthdayController) Delete(context echo.Context) error {
	birthdayID, err := c.GetParam(context, "birthday_id", "uint")
	if err != nil {
		return err
	}

	user, err := c.GetCurrentUser(context)
	if err != nil {
		return err
	}

	birthday, err := c.BirthdayService.GetOneByIdAndUser(birthdayID.(uint), user)
	if err != nil {
		return err
	}

	err = c.BirthdayService.Delete(birthday, user)
	if err != nil {
		return err
	}

	return c.Response(context, http.StatusOK, nil, "")
}
