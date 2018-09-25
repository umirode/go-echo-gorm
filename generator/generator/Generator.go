package generator

import (
    "github.com/umirode/go-rest/generator/generator/interfaces"
)

type Generator struct {
    FileCreator       interfaces.IFileCreator
    TemplateReader    interfaces.ITemplateReader
    TemplateGenerator interfaces.ITemplateGenerator
}

func (g *Generator) Generate(templateFilePath string, outputFilePath string, templateData interface{}) error {
    templateBytes, err := g.TemplateReader.Read(templateFilePath)
    if err != nil {
        return err
    }

    outputFile, err := g.FileCreator.Create(outputFilePath)
    defer outputFile.Close()
    if err != nil {
        return err
    }

    return g.TemplateGenerator.Generate(outputFile, templateBytes, templateData)
}
