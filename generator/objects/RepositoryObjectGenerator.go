package objects

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/umirode/go-rest/generator/generator/interfaces"
)

type RepositoryObjectGenerator struct {
	Generator interfaces.IGenerator
}

func (g *RepositoryObjectGenerator) Generate(name string, args []string) error {
	err := g.generateIRepository(name)
	if err != nil {
		return err
	}

	err = g.generateRepositoryMock(name)
	if err != nil {
		return err
	}

	err = g.generateDatabaseRepository(name)
	if err != nil {
		return err
	}

	return nil
}

func (g *RepositoryObjectGenerator) generateDatabaseRepository(name string) error {
	outputFilePath := fmt.Sprintf("../repositories/%sDatabaseRepository.go", strcase.ToCamel(name))

	return g.Generator.Generate("../repositories/templates/databaseRepository.txt", outputFilePath, struct {
		Name string
	}{
		Name: name,
	})
}

func (g *RepositoryObjectGenerator) generateRepositoryMock(name string) error {
	outputFilePath := fmt.Sprintf("../repositories/%sRepositoryMock.go", strcase.ToCamel(name))

	return g.Generator.Generate("../repositories/templates/repositoryMock.txt", outputFilePath, struct {
		Name string
	}{
		Name: name,
	})
}

func (g *RepositoryObjectGenerator) generateIRepository(name string) error {
	outputFilePath := fmt.Sprintf("../repositories/I%sRepository.go", strcase.ToCamel(name))

	return g.Generator.Generate("../repositories/templates/iRepository.txt", outputFilePath, struct {
		Name string
	}{
		Name: name,
	})
}
