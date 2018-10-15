package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/umirode/go-rest/configs"
	"github.com/umirode/go-rest/controllers"
	"github.com/umirode/go-rest/database"
	"github.com/umirode/go-rest/errors"
	"github.com/umirode/go-rest/repositories"
	"github.com/umirode/go-rest/services"
)

func main() {
	/**
	Load .env variables
	*/
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err.Error())
	}

	/**
	Create database connection
	*/
	databaseConfig := configs.GetDatabaseConfig()
	db, err := database.NewDatabase(
		&database.Config{
			Driver:   databaseConfig.Driver,
			Debug:    databaseConfig.Debug,
			Database: databaseConfig.Database,
			Host:     databaseConfig.Host,
			Port:     databaseConfig.Port,
			Username: databaseConfig.Username,
			Password: databaseConfig.Password,
			Params:   databaseConfig.Params,
		},
	)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	defer db.Close()

	/**
	Run migrations
	*/
	database.RunMigrations(db)

	/**
	Get server address
	*/
	serverConfig := configs.GetServerConfig()
	serverAddress := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)

	/**
	Get router
	*/
	router := getRouter(db)

	/**
	Start server
	*/
	logrus.Fatal(router.Start(serverAddress))
}

func getRouter(db *gorm.DB) *echo.Echo {
	router := echo.New()

	router.HTTPErrorHandler = errors.NewHTTPErrorHandler().Handler

	userController := &controllers.UserController{
		UserService: &services.UserService{
			Repository: repositories.NewUserDatabaseRepository(db),
		},
	}
	userGroup := router.Group("/users")
	userGroup.GET("", userController.GetAllUsers)
	userGroup.GET("/:id", userController.GetSingleUser)
	userGroup.POST("", userController.CreateUser)
	userGroup.PUT("/:id", userController.UpdateUser)
	userGroup.DELETE("/:id", userController.DeleteUser)

	return router
}
