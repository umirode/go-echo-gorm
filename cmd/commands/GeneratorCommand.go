package commands

import (
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/umirode/go-rest/generator/generator"
	"github.com/umirode/go-rest/generator/objects"
)

type GeneratorCommand struct {
	Generator *generator.Generator
}

func NewGeneratorCommand() *GeneratorCommand {
	return &GeneratorCommand{}
}

func (c *GeneratorCommand) WithDatabase(db *gorm.DB) ICommand {
	return c
}

func (c *GeneratorCommand) GetCommand() *cobra.Command {
	mainCommand := c.getMainCommand()

	basePath := mainCommand.Flags().StringP("path", "p", "", "base path for files")

	c.Generator = &generator.Generator{
		BasePath:          *basePath,
		FileCreator:       &generator.FileCreator{},
		TemplateReader:    &generator.TemplateReader{},
		TemplateGenerator: &generator.TemplateGenerator{},
	}

	mainCommand.AddCommand(c.getModelCommand())
	mainCommand.AddCommand(c.getRepositoryCommand())
	mainCommand.AddCommand(c.getConfigCommand())
	mainCommand.AddCommand(c.getMiddlewareCommand())
	mainCommand.AddCommand(c.getServiceCommand())

	return mainCommand
}

func (c *GeneratorCommand) getMainCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "generator",
		Short: "Boilerplate code generator",
		Long:  `Generate boilerplate code from templates`,
		Args:  cobra.MinimumNArgs(1),
	}

	return command
}

func (c *GeneratorCommand) getModelCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "model [name]",
		Short: "Generate model",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := objects.NewModelObjectGenerator(c.Generator).Generate(args[0], args)
			if err != nil {
				logrus.Fatal(err)
			}
		},
	}

	return command
}

func (c *GeneratorCommand) getRepositoryCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "repository [name]",
		Short: "Generate repository",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := objects.NewRepositoryObjectGenerator(c.Generator).Generate(args[0], args)
			if err != nil {
				logrus.Fatal(err)
			}
		},
	}

	return command
}

func (c *GeneratorCommand) getServiceCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "service [name]",
		Short: "Generate service",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := objects.NewServiceObjectGenerator(c.Generator).Generate(args[0], args)
			if err != nil {
				logrus.Fatal(err)
			}
		},
	}

	return command
}

func (c *GeneratorCommand) getConfigCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "config [name]",
		Short: "Generate config",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := objects.NewConfigObjectGenerator(c.Generator).Generate(args[0], args)
			if err != nil {
				logrus.Fatal(err)
			}
		},
	}

	return command
}

func (c *GeneratorCommand) getMiddlewareCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "middleware [name]",
		Short: "Generate middleware",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := objects.NewMiddlewareObjectGenerator(c.Generator).Generate(args[0], args)
			if err != nil {
				logrus.Fatal(err)
			}
		},
	}

	return command
}
