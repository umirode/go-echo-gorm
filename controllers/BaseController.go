package controllers

import (
	"github.com/umirode/go-rest/errors"
	"strconv"
)

type BaseController struct{}

type iGetParamContext interface {
	Param(name string) string
}

func (c *BaseController) getParam(context iGetParamContext, key string, valueType string) (interface{}, error) {
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
