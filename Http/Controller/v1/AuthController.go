package v1

import (
	"github.com/labstack/echo"
	"github.com/umirode/go-rest/Http/Controller"
	"github.com/umirode/go-rest/Http/Error"
	"github.com/umirode/go-rest/src/Domain/Service"
	"github.com/umirode/go-rest/src/Domain/Service/DTO"
	"net/http"
)

type AuthController struct {
	Controller.BaseController

	AuthService Service.IAuthService
}

func NewAuthController(authService Service.IAuthService, userService Service.IUserService) *AuthController {
	controller := &AuthController{
		AuthService: authService,
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

	accessToken, refreshToken, err := c.AuthService.Login(authDTO)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, struct {
		AccessToken     string `json:"access_token"`
		RefreshToken    string `json:"refresh_token"`
		AccessExpiresAt int64  `json:"access_expires_at"`
	}{
		AccessToken:     accessToken.Token,
		AccessExpiresAt: accessToken.ExpiresAt,
		RefreshToken:    refreshToken.Token,
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

func (c *AuthController) RefreshToken(context echo.Context) error {
	user, err := c.GetCurrentUser(context)
	if err != nil {
		return err
	}

	token, err := c.GetToken(context)
	if err != nil {
		return err
	}

	accessToken, refreshToken, err := c.AuthService.RefreshJWT(user, token.Raw)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, struct {
		AccessToken     string `json:"access_token"`
		RefreshToken    string `json:"refresh_token"`
		AccessExpiresAt int64  `json:"access_expires_at"`
	}{
		AccessToken:     accessToken.Token,
		AccessExpiresAt: accessToken.ExpiresAt,
		RefreshToken:    refreshToken.Token,
	})
}
