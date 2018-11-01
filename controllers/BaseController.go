package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/umirode/go-rest/errors"
	"strconv"
)

type BaseController struct {
	JWT struct {
		// Assess token
		ExpiresAt int64 // time in seconds
		Secret    string

		// Refresh token
		RefreshExpiresAt int64 // time in seconds
		RefreshSecret    string
	}
}

func (c *BaseController) getToken(context echo.Context) (*jwt.Token, error) {
	token, ok := context.Get("user").(*jwt.Token)
	if !ok {
		return nil, errors.NewInvalidTokenError()
	}

	return token, nil
}

func (c *BaseController) getTokenClaims(context echo.Context) (jwt.MapClaims, error) {
	token, err := c.getToken(context)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.NewInvalidTokenError()
	}

	return claims, nil
}

func (c *BaseController) getParam(context echo.Context, key string, valueType string) (interface{}, error) {
	param := context.Param(key)
	if param == "" {
		return nil, errors.NewRequestParsingError()
	}

	switch valueType {
	case "int":
		result, _ := strconv.Atoi(param)

		return result, nil
		break
	case "uint":
		result, _ := strconv.Atoi(param)

		return uint(result), nil
		break
	case "string":
		return param, nil
		break
	}

	return nil, errors.NewRequestParsingError()
}
