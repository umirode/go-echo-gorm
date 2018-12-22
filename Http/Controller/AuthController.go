package Controller

import (
	"github.com/labstack/echo"
	"github.com/umirode/go-rest/Http/Error"
	"github.com/umirode/go-rest/src/Domain/Service"
	"github.com/umirode/go-rest/src/Domain/Service/DTO"
	"net/http"
)

type AuthController struct {
	BaseController

	JWT struct {
		ExpiresAt int64 // time in seconds
		Secret    string
	}

	AuthService Service.IAuthService
}

func NewAuthController(authService Service.IAuthService, userService Service.IUserService, JWTExpiresAt int64, JWTSecret string) *AuthController {
	controller := &AuthController{
		AuthService: authService,
		JWT: struct {
			ExpiresAt int64
			Secret    string
		}{ExpiresAt: JWTExpiresAt, Secret: JWTSecret},
	}

	controller.UserService = userService

	return controller
}

func (c *AuthController) Login(context echo.Context) error {
	loginData := new(struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	})

	if err := context.Bind(loginData); err != nil {
		return Error.NewRequestParsingError()
	}

	if err := context.Validate(loginData); err != nil {
		return err
	}

	authDTO := &DTO.AuthDTO{
		Email:    loginData.Email,
		Password: loginData.Password,
	}

	user, err := c.AuthService.Login(authDTO)
	if err != nil {
		return err
	}

	jwtToken, err := c.AuthService.GetJWTTokenForUser(user, c.JWT.ExpiresAt, c.JWT.Secret)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, struct {
		Token     string `json:"token"`
		ExpiresAt int64  `json:"expires_at"`
	}{
		Token:     jwtToken.Token,
		ExpiresAt: jwtToken.ExpiresAt,
	})
}

func (c *AuthController) Signup(context echo.Context) error {
	signupData := new(struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	})

	if err := context.Bind(signupData); err != nil {
		return Error.NewRequestParsingError()
	}

	if err := context.Validate(signupData); err != nil {
		return err
	}

	authDTO := &DTO.AuthDTO{
		Email:    signupData.Email,
		Password: signupData.Password,
	}

	err := c.AuthService.Signup(authDTO)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, nil)
}

func (c *AuthController) ChangePassword(context echo.Context) error {
	resetPasswordData := new(struct {
		Password    string `json:"password" validate:"required"`
		NewPassword string `json:"new_password" validate:"required,min=8"`
	})

	if err := context.Bind(resetPasswordData); err != nil {
		return Error.NewRequestParsingError()
	}

	if err := context.Validate(resetPasswordData); err != nil {
		return err
	}

	user, err := c.getCurrentUser(context)
	if err != nil {
		return err
	}

	authDTO := &DTO.AuthDTO{
		Password:    resetPasswordData.Password,
		NewPassword: resetPasswordData.NewPassword,
	}

	err = c.AuthService.ChangePassword(user, authDTO)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, nil)
}
