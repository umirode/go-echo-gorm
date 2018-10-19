package generator

import (
	"fmt"
)

type Generator struct {
	BasePath string

	FileCreator       IFileCreator
	TemplateReader    ITemplateReader
	TemplateGenerator ITemplateGenerator
}

func (g *Generator) Generate(templateFilePath string, outputFilePath string, templateData interface{}) error {
	templateBytes, err := g.TemplateReader.Read(fmt.Sprintf("%s%s", g.BasePath, templateFilePath))
	if err != nil {
		return err
	}

	outputFile, err := g.FileCreator.Create(fmt.Sprintf("%s%s", g.BasePath, outputFilePath))
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return g.TemplateGenerator.Generate(outputFile, templateBytes, templateData)
}
