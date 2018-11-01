package router

import (
	"github.com/umirode/go-rest/configs"
	"github.com/umirode/go-rest/controllers"
	"github.com/umirode/go-rest/middlewares"
	"github.com/umirode/go-rest/repositories"
	"github.com/umirode/go-rest/services"
)

func (r *Router) setBirthdayRoutes() {
	config := configs.GetJWTConfig()

	birthdayController := &controllers.BirthdayController{
		BirthdayService: &services.BirthdayService{
			BirthdayRepository: repositories.NewBirthdayDatabaseRepository(r.Database),
		},
	}
	birthdayController.JWT.ExpiresAt = config.ExpiresAt
	birthdayController.JWT.Secret = config.Secret

	userBirthdaysGroup := r.Router.Group("/birthdays")
	userBirthdaysGroup.Use(middlewares.NewJWTMiddleware(config.Secret).Middleware)

	userBirthdaysGroup.GET("", birthdayController.GetAllBirthdaysForUser)
	userBirthdaysGroup.POST("", birthdayController.CreateBirthdayByUser)
	userBirthdaysGroup.PUT("/:birthday_id", birthdayController.UpdateBirthdayByUser)
	userBirthdaysGroup.DELETE("/:birthday_id", birthdayController.DeleteBirthdayByUser)
}
