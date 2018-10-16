package router

import (
	"github.com/umirode/go-rest/controllers"
	"github.com/umirode/go-rest/repositories"
	"github.com/umirode/go-rest/services"
)

func (r *Router) setUserRoutes() {
	userController := &controllers.UserController{
		UserService: &services.UserService{
			Repository: repositories.NewUserDatabaseRepository(r.Database),
		},
	}
	userGroup := r.Router.Group("/users")
	userGroup.GET("", userController.GetAllUsers)
	userGroup.GET("/:id", userController.GetSingleUser)
	userGroup.POST("", userController.CreateUser)
	userGroup.PUT("/:id", userController.UpdateUser)
	userGroup.DELETE("/:id", userController.DeleteUser)
}
