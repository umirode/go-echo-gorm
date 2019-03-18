package Router

import (
	"net/http"

	"github.com/iancoleman/strcase"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/umirode/go-rest/Module/Http/Middleware"
	"github.com/umirode/go-rest/Module/Http/Response"
	"github.com/umirode/go-rest/Validator"
	"github.com/umirode/go-rest/src/Common"
	goValidator "gopkg.in/go-playground/validator.v9"
)

type Router struct {
	Router *echo.Echo
	Debug  bool
}

func NewRouter(debug bool) *Router {
	router := &Router{
		Router: echo.New(),
		Debug:  debug,
	}
	router.Router.Validator = Validator.NewOnlyStructValidator()

	router.init()

	return router
}

func (r *Router) init() {
	if r.Debug {
		r.Router.Use(Middleware.NewLoggerMiddleware().Middleware)
	}

	r.Router.HTTPErrorHandler = NewHTTPErrorHandler().Handler

	r.Router.Use(Middleware.NewCorsMiddleware().Middleware)

	/**
	Set routes
	*/
	r.setV1Routes()
}

type HTTPErrorHandler struct{}

func NewHTTPErrorHandler() *HTTPErrorHandler {
	return &HTTPErrorHandler{}
}

func (h *HTTPErrorHandler) Handler(err error, context echo.Context) {
	response := new(Response.Response)

	switch v := err.(type) {
	case *echo.HTTPError:
		response.Status = v.Code
		response.Message = v.Message.(string)
		response.Data = nil
		break
	case Common.IHttpError:
		response.Status = v.Status()
		response.Message = v.Error()
		response.Data = nil
		break
	case goValidator.ValidationErrors:

		data := make(map[string][]string, 0)

		for _, validationErr := range v {
			field := strcase.ToSnake(validationErr.Field())
			data[field] = append(data[field], validationErr.Tag())
		}

		response.Status = http.StatusUnprocessableEntity
		response.Message = "Invalid"
		response.Data = data
		break
	default:
		response.Status = http.StatusInternalServerError
		response.Message = v.Error()
		response.Data = nil
		break
	}

	err = context.JSON(response.Status, response)
	if err != nil {
		logrus.Error(err)
	}
}
