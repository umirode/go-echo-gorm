package objects

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/umirode/go-rest/generator/generator"
)

type ModelObjectGenerator struct {
	Generator generator.IGenerator
}

func (g *ModelObjectGenerator) Generate(name string, args []string) error {
	err := g.generateModel(name)

	return err
}

func (g *ModelObjectGenerator) generateModel(name string) error {
	outputFilePath := fmt.Sprintf("models/%sModel.go", strcase.ToCamel(name))

	return g.Generator.Generate("models/templates/model.txt", outputFilePath, struct {
		Name string
	}{
		Name: name,
	})
}
