package commands

import (
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/umirode/go-rest/repositories"
	"github.com/umirode/go-rest/services"
)

type AuthCommand struct {
	Database *gorm.DB

	AuthService services.IAuthService
}

func NewAuthCommand() *AuthCommand {
	return &AuthCommand{}
}

func (c *AuthCommand) WithDatabase(db *gorm.DB) ICommand {
	c.Database = db

	return c
}

func (c *AuthCommand) GetCommand() *cobra.Command {
	mainCommand := c.getMainCommand()

	c.AuthService = &services.AuthService{
		UserRepository:            repositories.NewUserDatabaseRepository(c.Database),
		JWTRefreshTokenRepository: repositories.NewJWTRefreshTokenDatabaseRepository(c.Database),
	}

	mainCommand.AddCommand(c.getCreateCommand())
	mainCommand.AddCommand(c.getDeleteCommand())
	mainCommand.AddCommand(c.getDeleteRefreshTokensCommand())

	return mainCommand
}

func (c *AuthCommand) getMainCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "auth",
		Short: "Control users",
		Long:  `Create, delete users and reset user tokens`,
		Args:  cobra.MinimumNArgs(1),
	}

	return command
}

func (c *AuthCommand) getCreateCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "create [email] [password]",
		Short: "Create new user",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			err := c.AuthService.Signup(args[0], args[1])
			if err != nil {
				logrus.Fatal(err)
			}
		},
	}

	return command
}

func (c *AuthCommand) getDeleteCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "delete [email]",
		Short: "Delete user",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := c.AuthService.DeleteUserByEmail(args[0])
			if err != nil {
				logrus.Fatal(err)
			}
		},
	}

	return command
}

func (c *AuthCommand) getDeleteRefreshTokensCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "delete-tokens [email]",
		Short: "Generate model",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := c.AuthService.DeleteUserRefreshTokensIfMoreByEmail(args[0], 0)
			if err != nil {
				logrus.Fatal(err)
			}
		},
	}

	return command
}
