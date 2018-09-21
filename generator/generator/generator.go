package generator

import (
	"github.com/iancoleman/strcase"
	"io"
	"io/ioutil"
	"os"
	"text/template"
)

type IFileCreator interface {
	Create(filePath string) (*os.File, error)
}

type FileCreator struct{}

func (fc *FileCreator) Create(filePath string) (*os.File, error) {
	newFile, err := os.Create(filePath)
	if err != nil {
		os.Remove(filePath)
		return &os.File{}, err
	}

	return newFile, nil
}

type ITemplateReader interface {
	Read(filePath string) ([]byte, error)
}

type TemplateReader struct{}

func (tr *TemplateReader) Read(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

type ITemplateGenerator interface {
	Generate(file io.Writer, templateBytes []byte, templateData interface{}) error
}

type TemplateGenerator struct{}

func (tg *TemplateGenerator) Generate(file io.Writer, templateBytes []byte, templateData interface{}) error {
	fileTemplate := template.Must(template.New("").Funcs(template.FuncMap{
		"ToCamel":      strcase.ToCamel,
		"ToLowerCamel": strcase.ToLowerCamel,
	}).Parse(string(templateBytes)))

	return fileTemplate.Execute(file, templateData)
}

type IGenerator interface {
	Generate(templateFilePath string, outputFilePath string, templateData interface{}) error
}

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
	defer outputFile.Close()
	if err != nil {
		return err
	}

	return g.TemplateGenerator.Generate(outputFile, templateBytes, templateData)
}
