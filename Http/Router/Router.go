package Router

import (
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
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
	message := new(struct {
		Error interface{} `json:"error"`
	})

	switch v := err.(type) {
	case *echo.HTTPError:
		message.Error = v.Message
		context.JSON(v.Code, message)
		break
	case Common.IHttpError:
		message.Error = v.Error()
		context.JSON(v.Status(), message)
		break
	case goValidator.ValidationErrors:

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
