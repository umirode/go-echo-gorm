package Router

import (
	"github.com/umirode/go-rest/Config"
	"github.com/umirode/go-rest/Http/Controller/v1"
	"github.com/umirode/go-rest/Http/Middleware"
	"github.com/umirode/go-rest/src/Application/Service"
	"github.com/umirode/go-rest/src/Infrastructure/Persistence/Gorm/Repository"
)

func (r *Router) setV1Routes() {
	config := Config.GetJWTConfig()

	userRepository := Repository.NewUserRepository(r.Database)
	birthdayRepository := Repository.NewBirthdayRepository(r.Database)
	refreshTokenRepository := Repository.NewRefreshTokenRepository(r.Database)
	notificationTokenRepository := Repository.NewNotificationTokenRepository(r.Database)

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

	notificationController := v1.NewNotificationController(
		Service.NewNotificationTokenService(notificationTokenRepository),
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

	notificationRoutes := v1Routes.Group("/notifications")
	notificationRoutes.Use(Middleware.NewJWTMiddleware(config.AccessTokenSecret).Middleware)

	notificationRoutes.POST("/tokens", notificationController.SaveToken)
}
