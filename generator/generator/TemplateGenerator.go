package generator

import (
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"io"
	"text/template"
)

type TemplateGenerator struct{}

func (tg *TemplateGenerator) Generate(file io.Writer, templateBytes []byte, templateData interface{}) error {
	fileTemplate := template.Must(
		template.
			New("").
			Funcs(template.FuncMap{
				"ToCamel":      strcase.ToCamel,
				"ToLowerCamel": strcase.ToLowerCamel,
				"ToSnake":      strcase.ToSnake,
				"Plural":       inflection.Plural,
				"Singular":     inflection.Singular,
			}).
			Parse(string(templateBytes)),
	)

	return fileTemplate.Execute(file, templateData)
}
