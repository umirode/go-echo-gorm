package objects

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/umirode/go-rest/generator/generator/interfaces"
)

type ConfigObjectGenerator struct {
	Generator interfaces.IGenerator
}

func (g *ConfigObjectGenerator) Generate(name string, args []string) error {
	err := g.generateConfig(name)

	return err
}

func (g *ConfigObjectGenerator) generateConfig(name string) error {
	outputFilePath := fmt.Sprintf("configs/%sConfig.go", strcase.ToCamel(name))

	return g.Generator.Generate("configs/templates/config.txt", outputFilePath, struct {
		Name string
	}{
		Name: name,
	})
}
