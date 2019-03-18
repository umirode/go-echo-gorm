package Router

import (
	"github.com/sirupsen/logrus"
	"github.com/umirode/go-rest/Config"
	"github.com/umirode/go-rest/Module/Http/Controller/v1"
	"github.com/umirode/go-rest/Module/Http/Middleware"
	"github.com/umirode/go-rest/src/Application/Service"
	"github.com/umirode/go-rest/src/Infrastructure/Persistence/Gorm/Repository"
	"github.com/umirode/golang-echo-socket.io"
)

func (r *Router) setV1Routes() {
	config := Config.GetJWTConfig()

	userRepository := Repository.NewUserRepository()
	birthdayRepository := Repository.NewBirthdayRepository()
	refreshTokenRepository := Repository.NewRefreshTokenRepository()
	notificationTokenRepository := Repository.NewNotificationTokenRepository()

	authController := v1.NewAuthController(
		Service.NewAuthService(
			userRepository,
		),
		Service.NewJWTAuthService(
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

	v1Routes.Any("/socket.io/", func() *golang_echo_socket_io.Wrapper {
		socketIOController := v1.NewSocketIOController()

		wrapper, err := golang_echo_socket_io.NewWrapper(nil)
		if err != nil {
			logrus.Error(err)
		}

		wrapper.OnConnect("/", socketIOController.OnConnect)
		wrapper.OnError("/", socketIOController.OnError)
		wrapper.OnDisconnect("/", socketIOController.OnDisconnect)

		wrapper.OnEvent("/", "test", socketIOController.OnEventTest)

		go wrapper.Server.Serve()

		return wrapper
	}().HandlerFunc)
}
