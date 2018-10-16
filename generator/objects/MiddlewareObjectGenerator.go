package objects

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/umirode/go-rest/generator/generator/interfaces"
)

type MiddlewareObjectGenerator struct {
	Generator interfaces.IGenerator
}

func (g *MiddlewareObjectGenerator) Generate(name string, args []string) error {
	err := g.generateMiddleware(name)

	return err
}

func (g *MiddlewareObjectGenerator) generateMiddleware(name string) error {
	outputFilePath := fmt.Sprintf("middlewares/%sMiddleware.go", strcase.ToCamel(name))

	return g.Generator.Generate("middlewares/templates/middleware.txt", outputFilePath, struct {
		Name string
	}{
		Name: name,
	})
}
