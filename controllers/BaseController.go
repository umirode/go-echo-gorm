package controllers

import (
	"errors"
	"strconv"
)

type BaseController struct{}

type iGetParamContext interface {
	Param(name string) string
}

func (c *BaseController) getParam(context iGetParamContext, key string, valueType string) interface{} {
	param := context.Param(key)
	if param == "" {
		return nil
	}

	switch valueType {
	case "int":
		result, _ := strconv.Atoi(param)

		return result
		break
	case "uint":
		result, _ := strconv.Atoi(param)

		return uint(result)
		break
	case "string":
		return param
		break
	}

	return nil
}

func (c *BaseController) getNotExistsError() error {
	return errors.New("Not exists ")
}

func (c *BaseController) getQueryParsingError() error {
	return errors.New("Query parsing error ")
}
