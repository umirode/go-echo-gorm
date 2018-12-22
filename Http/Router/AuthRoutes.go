package Router

import (
	"github.com/umirode/go-rest/Http/Controller"
	"github.com/umirode/go-rest/Http/Middleware"
	"github.com/umirode/go-rest/configs"
	"github.com/umirode/go-rest/src/Application/Service"
	"github.com/umirode/go-rest/src/Infrastructure/Persistence/Gorm/Repository"
)

func (r *Router) setAuthRoutes() {
	config := configs.GetJWTConfig()

	userRepository := Repository.NewUserRepository(r.Database)

	authController := Controller.NewAuthController(Service.NewAuthService(userRepository), Service.NewUserService(userRepository), config.ExpiresAt, config.Secret)

	authGroup := r.Router.Group("/auth")

	authGroup.POST("/login", authController.Login)
	authGroup.POST("/signup", authController.Signup)

	protectedAuthGroup := authGroup.Group("")
	protectedAuthGroup.Use(Middleware.NewJWTMiddleware(config.Secret).Middleware)

	protectedAuthGroup.POST("/change-password", authController.ChangePassword)
}
