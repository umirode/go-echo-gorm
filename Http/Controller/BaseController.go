package Controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/umirode/go-rest/Http/Error"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Domain/Service"
	"strconv"
)

type BaseController struct {
	UserService Service.IUserService
}

func (c *BaseController) GetToken(context echo.Context) (*jwt.Token, error) {
	token, ok := context.Get("user").(*jwt.Token)
	if !ok {
		return nil, Error.NewInvalidTokenError()
	}

	return token, nil
}

func (c *BaseController) GetTokenClaims(context echo.Context) (jwt.MapClaims, error) {
	token, err := c.GetToken(context)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, Error.NewInvalidTokenError()
	}

	return claims, nil
}

func (c *BaseController) GetParam(context echo.Context, key string, valueType string) (interface{}, error) {
	param := context.Param(key)
	if param == "" {
		return nil, Error.NewRequestParsingError()
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

	return nil, Error.NewRequestParsingError()
}

func (c *BaseController) GetCurrentUser(context echo.Context) (*Entity.User, error) {
	claims, err := c.GetTokenClaims(context)
	if err != nil {
		return nil, err
	}
	userID := uint(claims["user_id"].(float64))

	return c.UserService.GetOneById(userID)
}
