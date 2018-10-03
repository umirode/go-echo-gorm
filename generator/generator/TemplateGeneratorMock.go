package generator

import (
	"github.com/stretchr/testify/mock"
	"io"
)

type TemplateGeneratorMock struct {
	mock.Mock
}

func (tg *TemplateGeneratorMock) Generate(file io.Writer, templateBytes []byte, templateData interface{}) error {
	args := tg.Called(file, templateBytes, templateData)

	return args.Error(0)
}
