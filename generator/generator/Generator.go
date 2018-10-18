package generator

type Generator struct {
	FileCreator       IFileCreator
	TemplateReader    ITemplateReader
	TemplateGenerator ITemplateGenerator
}

func (g *Generator) Generate(templateFilePath string, outputFilePath string, templateData interface{}) error {
	templateBytes, err := g.TemplateReader.Read(templateFilePath)
	if err != nil {
		return err
	}

	outputFile, err := g.FileCreator.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return g.TemplateGenerator.Generate(outputFile, templateBytes, templateData)
}
