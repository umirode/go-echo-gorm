package router

import (
	"github.com/umirode/go-rest/configs"
	"github.com/umirode/go-rest/controllers"
	"github.com/umirode/go-rest/middlewares"
	"github.com/umirode/go-rest/repositories"
	"github.com/umirode/go-rest/services"
)

func (r *Router) setUserRoutes() {
	config := configs.GetJWTConfig()

	userController := &controllers.UserController{
		UserService: &services.UserService{
			UserRepository: repositories.NewUserDatabaseRepository(r.Database),
		},
	}

	userGroup := r.Router.Group("/users")

	userGroup.Use(middlewares.NewJWTMiddleware(config.Secret).Middleware)

	userGroup.GET("", userController.GetAllUsers)
	userGroup.GET("/:id", userController.GetSingleUser)
	userGroup.POST("", userController.CreateUser)
	userGroup.PUT("/:id", userController.UpdateUser)
	userGroup.DELETE("/:id", userController.DeleteUser)
}
