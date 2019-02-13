package v1

import (
	"github.com/labstack/echo"
	"github.com/umirode/go-rest/Http/Controller"
	"github.com/umirode/go-rest/Http/Error"
	"github.com/umirode/go-rest/src/Domain/Service"
	"github.com/umirode/go-rest/src/Domain/Service/DTO"
	"net/http"
)

type NotificationController struct {
	Controller.BaseController

	NotificationTokenService Service.INotificationTokenService
}

func NewNotificationController(notificationTokenService Service.INotificationTokenService, userService Service.IUserService) *NotificationController {
	controller := &NotificationController{
		NotificationTokenService: notificationTokenService,
	}

	controller.UserService = userService

	return controller
}

func (c *NotificationController) SaveToken(context echo.Context) error {
	data := new(struct {
		Token string `json:"name" validate:"required,min=10,max=255"`
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

	notificationTokenDTO := &DTO.NotificationTokenDTO{
		Token: data.Token,
	}

	err = c.NotificationTokenService.Create(notificationTokenDTO, user)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, nil)
}
