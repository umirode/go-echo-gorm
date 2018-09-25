package main

import (
    "fmt"
    "github.com/Sirupsen/logrus"
    "github.com/jinzhu/gorm"
    "github.com/labstack/echo"
    "github.com/umirode/go-rest/configs"
    "github.com/umirode/go-rest/controllers"
    "github.com/umirode/go-rest/database"
    "github.com/umirode/go-rest/models"
    "github.com/umirode/go-rest/repositories"
    "github.com/umirode/go-rest/response"
    "github.com/umirode/go-rest/services"
)

func main() {
    databaseConfig := configs.GetDatabaseConfig()
    db, _ := database.NewDatabase(
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

    runMigrations(db)
    r := getRouter(db)

    serverConfig := configs.GetServerConfig()
    serverAddress := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
    logrus.Fatal(r.Start(serverAddress))
}

func getRouter(db *gorm.DB) *echo.Echo {
    router := echo.New()

    router.HTTPErrorHandler = func(err error, context echo.Context) {
        response.SendResponseJson(context, "error", err.Error(), "")
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

func runMigrations(db *gorm.DB) {
    db.AutoMigrate(
        &models.UserModel{},
    )
}
