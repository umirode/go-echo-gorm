package Router

import (
	"github.com/umirode/go-rest/Config"
	"github.com/umirode/go-rest/Http/Controller/v1"
	"github.com/umirode/go-rest/Http/Middleware"
	"github.com/umirode/go-rest/src/Application/Service"
	"github.com/umirode/go-rest/src/Infrastructure/Persistence/Gorm/Repository"
)

/**
POST /v1/auth/login
POST /v1/auth/signup
POST /v1/auth/refresh-token
POST /v1/auth/logout

GET /v1/birthdays
POST /v1/birthdays
PUT /v1/birthdays/{id}
GET /v1/birthdays/{id}
DELETE /v1/birthdays/{id}
*/
func (r *Router) setV1Routes() {
	config := Config.GetJWTConfig()

	userRepository := Repository.NewUserRepository(r.Database)
	birthdayRepository := Repository.NewBirthdayRepository(r.Database)
	refreshTokenRepository := Repository.NewRefreshTokenRepository(r.Database)

	authController := v1.NewAuthController(
		Service.NewAuthService(
			userRepository,
			refreshTokenRepository,
			config.AccessTokenSecret,
			config.AccessTokenLifeTime,
			config.RefreshTokenSecret,
			config.RefreshTokenLifeTime,
		),
		Service.NewUserService(userRepository),
	)

	birthdayController := v1.NewBirthdayController(
		Service.NewBirthdayService(birthdayRepository),
		Service.NewUserService(userRepository),
	)

	v1Routes := r.Router.Group("/v1")

	authRoutes := v1Routes.Group("/auth")
	authRoutes.POST("/login", authController.Login)
	authRoutes.POST("/signup", authController.Signup)
	authRoutes.POST("/refresh-token", authController.RefreshToken, Middleware.NewJWTMiddleware(config.RefreshTokenSecret).Middleware)

	birthdayRoutes := v1Routes.Group("/birthdays")
	birthdayRoutes.Use(Middleware.NewJWTMiddleware(config.AccessTokenSecret).Middleware)

	birthdayRoutes.GET("", birthdayController.GetAll)
	birthdayRoutes.POST("", birthdayController.Create)

	birthdayRoutes.GET("/:birthday_id", birthdayController.GetOne)
	birthdayRoutes.PUT("/:birthday_id", birthdayController.Update)
	birthdayRoutes.DELETE("/:birthday_id", birthdayController.Delete)
}
