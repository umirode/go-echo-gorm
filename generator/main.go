package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/inflection"
	"github.com/pkg/errors"
	"github.com/umirode/go-rest/generator/generator"
	"github.com/umirode/go-rest/generator/objects"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "Boilerplate code generator"
	app.Usage = ""
	app.UsageText = "GENERATOR -t TYPE -n NAME -p PARAMS"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "type, t",
			Usage: "What are you want to generate",
		},
		&cli.StringFlag{
			Name:  "name, n",
			Usage: "Name for generated object",
		},
		&cli.StringSliceFlag{
			Name:  "param, p",
			Usage: "Params",
		},
	}

	app.Action = func(c *cli.Context) error {
		object := inflection.Singular(c.String("type"))
		name := inflection.Singular(c.String("name"))
		params := c.StringSlice("param")

		if name == "" {
			return errors.New("Name can not be blank")
		}

		if object == "" {
			return errors.New("Type can not be blank")
		}

		handler := getHandlers()[object]
		if handler == nil {
			return errors.New(fmt.Sprintf("Type \"%s\" not found", object))
		}

		err := handler.Generate(name, params)

		return err
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getHandlers() map[string]objects.IObjectGenerator {
	gen := &generator.Generator{
		FileCreator:       &generator.FileCreator{},
		TemplateReader:    &generator.TemplateReader{},
		TemplateGenerator: &generator.TemplateGenerator{},
	}

	handlers := map[string]objects.IObjectGenerator{
		"model": &objects.ModelObjectGenerator{
			Generator: gen,
		},
		"repository": &objects.RepositoryObjectGenerator{
			Generator: gen,
		},
		"service": &objects.ServiceObjectGenerator{
			Generator: gen,
		},
		"config": &objects.ConfigObjectGenerator{
			Generator: gen,
		},
		"middleware": &objects.MiddlewareObjectGenerator{
			Generator: gen,
		},
	}

	return handlers
}
