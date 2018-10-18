package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/umirode/go-rest/errors"
	"github.com/umirode/go-rest/services"
	"net/http"
)

type AuthController struct {
	BaseController

	AuthService services.IAuthService

	JWT struct {
		// Assess token
		ExpiresAt int64 // time in seconds
		Secret    string

		// Refresh token
		RefreshExpiresAt int64 // time in seconds
		RefreshSecret    string
	}
}

func (c *AuthController) Login(context echo.Context) error {
	loginData := new(struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	})

	if err := context.Bind(loginData); err != nil {
		return errors.NewRequestParsingError()
	}

	if err := context.Validate(loginData); err != nil {
		return err
	}

	token, refreshToken, expiresAt, err := c.AuthService.Login(loginData.Email, loginData.Password, context.RealIP(), services.JWTConfig{
		ExpiresAt:        c.JWT.ExpiresAt,
		Secret:           c.JWT.Secret,
		RefreshExpiresAt: c.JWT.RefreshExpiresAt,
		RefreshSecret:    c.JWT.RefreshSecret,
	})
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresAt    int64  `json:"expires_at"`
	}{
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	})
}

func (c *AuthController) Signup(context echo.Context) error {
	signupData := new(struct {
		Email                string `json:"email" validate:"required,email"`
		Password             string `json:"password" validate:"required,eqfield=PasswordConfirmation"`
		PasswordConfirmation string `json:"password_confirmation" validate:"required"`
	})

	if err := context.Bind(signupData); err != nil {
		return errors.NewRequestParsingError()
	}

	if err := context.Validate(signupData); err != nil {
		return err
	}

	err := c.AuthService.Signup(signupData.Email, signupData.Password)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return context.JSON(http.StatusOK, nil)
}

func (c *AuthController) RefreshToken(context echo.Context) error {
	token := context.Get("user").(*jwt.Token)

	assessToken, refreshToken, expiresAt, err := c.AuthService.RefreshToken(token, services.JWTConfig{
		ExpiresAt:        c.JWT.ExpiresAt,
		Secret:           c.JWT.Secret,
		RefreshExpiresAt: c.JWT.RefreshExpiresAt,
		RefreshSecret:    c.JWT.RefreshSecret,
	})
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresAt    int64  `json:"expires_at"`
	}{
		Token:        assessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	})
}

func (c *AuthController) Logout(context echo.Context) error {
	token := context.Get("user").(*jwt.Token)
	err := c.AuthService.Logout(token)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, nil)
}

func (c *AuthController) ResetPassword(context echo.Context) error {
	resetPasswordData := new(struct {
		Password                string `json:"password" validate:"required"`
		NewPassword             string `json:"new_password" validate:"required,eqfield=NewPasswordConfirmation"`
		NewPasswordConfirmation string `json:"new_password_confirmation" validate:"required"`
	})

	if err := context.Bind(resetPasswordData); err != nil {
		return errors.NewRequestParsingError()
	}

	if err := context.Validate(resetPasswordData); err != nil {
		return err
	}

	token := context.Get("user").(*jwt.Token)

	err := c.AuthService.ResetPassword(
		token,
		resetPasswordData.Password,
		resetPasswordData.NewPassword,
	)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, nil)
}
