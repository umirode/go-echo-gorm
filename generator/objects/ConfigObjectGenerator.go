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
	if err != nil {
		return err
	}

	err = g.generateConfigTest(name)
	if err != nil {
		return err
	}

	return nil
}

func (g *ConfigObjectGenerator) generateConfig(name string) error {
	outputFilePath := fmt.Sprintf("../configs/%s.go", strcase.ToCamel(name))

	return g.Generator.Generate("../configs/templates/config.txt", outputFilePath, struct {
		Name string
	}{
		Name: name,
	})
}

func (g *ConfigObjectGenerator) generateConfigTest(name string) error {
	outputFilePath := fmt.Sprintf("../configs/%s_test.go", strcase.ToCamel(name))

	return g.Generator.Generate("../configs/templates/configTest.txt", outputFilePath, struct {
		Name string
	}{
		Name: name,
	})
}
