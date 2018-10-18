package errors

import (
	"github.com/iancoleman/strcase"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type HTTPErrorHandler struct{}

func NewHTTPErrorHandler() *HTTPErrorHandler {
	return &HTTPErrorHandler{}
}

func (h *HTTPErrorHandler) Handler(err error, context echo.Context) {
	message := new(struct {
		Error interface{} `json:"error"`
	})

	switch v := err.(type) {
	case *echo.HTTPError:
		message.Error = v.Message
		context.JSON(v.Code, message)
		break
	case IHTTPError:
		message.Error = v.Error()
		context.JSON(v.Status(), message)
		break
	case validator.ValidationErrors:

		data := make(map[string][]string, 0)

		for _, validationErr := range v {
			field := strcase.ToSnake(validationErr.Field())
			data[field] = append(data[field], validationErr.Tag())
		}

		message.Error = data
		context.JSON(http.StatusUnprocessableEntity, message)
		break
	default:
		message.Error = v.Error()
		context.JSON(http.StatusInternalServerError, message)
		break
	}
}
