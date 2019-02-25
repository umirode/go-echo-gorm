package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/umirode/go-rest/Cli/Command"
)

func main() {
	/**
	Load .env variables
	*/
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err)
	}

	rootCmd := &cobra.Command{Use: "cmd"}

	c := getCommands()
	for _, command := range c {
		rootCmd.AddCommand(
			command.GetCommand(),
		)
	}

	err = rootCmd.Execute()
	if err != nil {
		logrus.Fatal(err)
	}
}

func getCommands() []Command.ICommand {

	return []Command.ICommand{}
}
