package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/umirode/go-rest/configs"
	"github.com/umirode/go-rest/database"
	"github.com/umirode/go-rest/models"
)

func main() {
	databaseConfig := configs.GetDatabaseConfig()
	database.OpenDBConnection(
		databaseConfig.Driver,
		databaseConfig.Database,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Params,
	)
	defer database.CloseDBConnection()

	runMigrations()

	serverConfig := configs.GetServerConfig()
	serverAddress := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	router := getRouter()
	logrus.Fatal(router.Start(serverAddress))
}

func getRouter() *echo.Echo {
	e := echo.New()

	return e
}

func runMigrations() {
	dbConnection := database.GetDBConnection()

	if dbConnection != nil {
		dbConnection.AutoMigrate(
			&models.UserModel{},
		)
	}
}
