package Router

import (
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/umirode/go-rest/Http"
	"github.com/umirode/go-rest/Http/Middleware"
	"github.com/umirode/go-rest/Validator"
	"github.com/umirode/go-rest/src/Common"
	goValidator "gopkg.in/go-playground/validator.v9"
	"net/http"
)

type Router struct {
	Router   *echo.Echo
	Database *gorm.DB
	Debug    bool
}

func NewRouter(database *gorm.DB, debug bool) *Router {
	router := &Router{
		Router:   echo.New(),
		Database: database,
		Debug:    debug,
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

	r.setV1Routes()
}

type HTTPErrorHandler struct{}

func NewHTTPErrorHandler() *HTTPErrorHandler {
	return &HTTPErrorHandler{}
}

func (h *HTTPErrorHandler) Handler(err error, context echo.Context) {
	response := new(Http.Response)

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

	context.JSON(response.Status, response)
}
