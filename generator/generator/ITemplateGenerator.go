package generator

import (
	"io"
)

type ITemplateGenerator interface {
	Generate(file io.Writer, templateBytes []byte, templateData interface{}) error
}
