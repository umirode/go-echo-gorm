package router

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/umirode/go-rest/errors"
	"github.com/umirode/go-rest/middlewares"
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

	router.init()

	return router
}

func (r *Router) init() {
	if r.Debug {
		r.Router.Use(middlewares.NewLoggerMiddleware().Middleware)
	}

	r.Router.HTTPErrorHandler = errors.NewHTTPErrorHandler().Handler

	r.setUserRoutes()
}
