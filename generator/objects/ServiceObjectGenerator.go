package objects

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/umirode/go-rest/generator/generator/interfaces"
)

type ServiceObjectGenerator struct {
	Generator interfaces.IGenerator
}

func (g *ServiceObjectGenerator) Generate(name string, args []string) error {
	err := g.generateIService(name)
	if err != nil {
		return err
	}

	err = g.generateServiceMock(name)
	if err != nil {
		return err
	}

	err = g.generateService(name)
	if err != nil {
		return err
	}

	return nil
}

func (g *ServiceObjectGenerator) generateService(name string) error {
	outputFilePath := fmt.Sprintf("services/%sService.go", strcase.ToCamel(name))

	return g.Generator.Generate("services/templates/service.txt", outputFilePath, struct {
		Name string
	}{
		Name: name,
	})
}

func (g *ServiceObjectGenerator) generateServiceMock(name string) error {
	outputFilePath := fmt.Sprintf("services/%sServiceMock.go", strcase.ToCamel(name))

	return g.Generator.Generate("services/templates/serviceMock.txt", outputFilePath, struct {
		Name string
	}{
		Name: name,
	})
}

func (g *ServiceObjectGenerator) generateIService(name string) error {
	outputFilePath := fmt.Sprintf("services/I%sService.go", strcase.ToCamel(name))

	return g.Generator.Generate("services/templates/iService.txt", outputFilePath, struct {
		Name string
	}{
		Name: name,
	})
}
