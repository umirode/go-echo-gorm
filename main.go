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
	"github.com/umirode/go-rest/repositories"
	"github.com/umirode/go-rest/response"
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
			Database: databaseConfig.Database,
			Host:     databaseConfig.Host,
			Port:     databaseConfig.Port,
			Username: databaseConfig.Username,
			Password: databaseConfig.Password,
			Params:   databaseConfig.Params,
		},
	)
	defer db.Close()
	if err != nil {
		db.Close()
		logrus.Fatal(err.Error())
	}

	/**
	Run migrations
	*/
	database.RunMigrations(db)

	/**
	Get router
	*/
	router := getRouter(db)

	/**
	Get server address
	*/
	serverConfig := configs.GetServerConfig()
	serverAddress := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)

	/**
	Start server
	*/
	logrus.Fatal(router.Start(serverAddress))
}

func getRouter(db *gorm.DB) *echo.Echo {
	router := echo.New()

	router.HTTPErrorHandler = func(err error, context echo.Context) {
		message := ""

		httpErr, httpErrOk := err.(*echo.HTTPError)
		if httpErrOk {
			message = httpErr.Message.(string)
		} else {
			message = err.Error()
		}

		response.SendResponseJson(context, "error", message, "")
	}

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
