package generator

type IGenerator interface {
	Generate(templateFilePath string, outputFilePath string, templateData interface{}) error
}
