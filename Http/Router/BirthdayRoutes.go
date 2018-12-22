package Router

import (
	"github.com/umirode/go-rest/Http/Controller"
	"github.com/umirode/go-rest/Http/Middleware"
	"github.com/umirode/go-rest/configs"
	"github.com/umirode/go-rest/src/Application/Service"
	"github.com/umirode/go-rest/src/Infrastructure/Persistence/Gorm/Repository"
)

func (r *Router) setBirthdayRoutes() {
	config := configs.GetJWTConfig()

	userRepository := Repository.NewUserRepository(r.Database)
	birthdayRepository := Repository.NewBirthdayRepository(r.Database)

	birthdayController := Controller.NewBirthdayController(Service.NewBirthdayService(birthdayRepository), Service.NewUserService(userRepository))

	userBirthdaysGroup := r.Router.Group("/birthdays")
	userBirthdaysGroup.Use(Middleware.NewJWTMiddleware(config.Secret).Middleware)

	userBirthdaysGroup.GET("", birthdayController.Index)
	userBirthdaysGroup.POST("", birthdayController.Create)
	userBirthdaysGroup.PUT("/:birthday_id", birthdayController.Update)
	userBirthdaysGroup.DELETE("/:birthday_id", birthdayController.Delete)
}
