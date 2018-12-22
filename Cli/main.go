package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/umirode/go-rest/Cli/commands"
	"github.com/umirode/go-rest/Database"
	"github.com/umirode/go-rest/configs"
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

func getCommands(db *gorm.DB) []commands.ICommand {
	return []commands.ICommand{}
}
