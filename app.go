package main

import (
    "github.com/umirode/go-rest/configs"
    "github.com/umirode/go-rest/models"
    "fmt"
    "github.com/Sirupsen/logrus"
    "github.com/labstack/echo"
)

func main() {
    databaseConfig := configs.GetDatabaseConfig()
    models.OpenDBConnection(
        databaseConfig.Driver,
        databaseConfig.Database,
        databaseConfig.Host,
        databaseConfig.Port,
        databaseConfig.Username,
        databaseConfig.Password,
        databaseConfig.Params,
    )
    defer models.CloseDBConnection()

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
    dbConnection := models.GetDBConnection()

    if dbConnection != nil {
        dbConnection.AutoMigrate()
    }
}
