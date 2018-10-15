package errors

import (
	"github.com/labstack/echo"
	"github.com/umirode/go-rest/response"
	"net/http"
)

type HTTPErrorHandler struct{}

func NewHTTPErrorHandler() *HTTPErrorHandler {
	return &HTTPErrorHandler{}
}

func (h *HTTPErrorHandler) Handler(err error, context echo.Context) {
	status := http.StatusInternalServerError

	switch v := err.(type) {
	case *echo.HTTPError:
		status = v.Code
		break
	case *AlreadyExistsError:
		status = v.Status
		break
	case *NotFoundError:
		status = v.Status
		break
	case *RequestParsingError:
		status = v.Status
		break
	}

	response.SendResponseJson(context, status, nil)
}
