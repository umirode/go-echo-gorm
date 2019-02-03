package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/umirode/go-rest/Cli/Command"
	"github.com/umirode/go-rest/Config"
	"github.com/umirode/go-rest/Database"
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
	databaseConfig := Config.GetDatabaseConfig()
	db, err := Database.NewDatabase(
		&Database.Config{
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
	Database.RunMigrations(db)

	rootCmd := &cobra.Command{Use: "cmd"}

	c := getCommands(db)
	for _, command := range c {
		rootCmd.AddCommand(
			command.GetCommand(),
		)
	}

	rootCmd.Execute()
}

func getCommands(db *gorm.DB) []Command.ICommand {
	return []Command.ICommand{}
}
