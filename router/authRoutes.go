package router

import (
	"github.com/umirode/go-rest/configs"
	"github.com/umirode/go-rest/controllers"
	"github.com/umirode/go-rest/middlewares"
	"github.com/umirode/go-rest/repositories"
	"github.com/umirode/go-rest/services"
)

func (r *Router) setAuthRoutes() {
	config := configs.GetJWTConfig()

	authController := &controllers.AuthController{
		JWT: struct {
			ExpiresAt        int64
			Secret           string
			RefreshExpiresAt int64
			RefreshSecret    string
		}{
			ExpiresAt:        config.ExpiresAt,
			Secret:           config.Secret,
			RefreshExpiresAt: config.RefreshExpiresAt,
			RefreshSecret:    config.RefreshSecret,
		},
		AuthService: &services.AuthService{
			UserRepository:            repositories.NewUserDatabaseRepository(r.Database),
			JWTRefreshTokenRepository: repositories.NewJWTRefreshTokenDatabaseRepository(r.Database),
		},
	}

	authGroup := r.Router.Group("/auth")

	authGroup.POST("/login", authController.Login)
	authGroup.POST("/signup", authController.Signup)

	authGroup.POST("/refresh-token", authController.RefreshToken, middlewares.NewJWTMiddleware(config.RefreshSecret).Middleware)

	protectedAuthGroup := authGroup.Group("")
	protectedAuthGroup.Use(middlewares.NewJWTMiddleware(config.Secret).Middleware)

	protectedAuthGroup.POST("/logout", authController.Logout)
	protectedAuthGroup.POST("/reset-password", authController.ResetPassword)
}
