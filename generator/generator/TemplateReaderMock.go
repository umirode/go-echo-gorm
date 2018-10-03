package generator

import (
	"github.com/stretchr/testify/mock"
)

type TemplateReaderMock struct {
	mock.Mock
}

func (tr *TemplateReaderMock) Read(filePath string) ([]byte, error) {
	args := tr.Called(filePath)

	return args.Get(0).([]byte), args.Error(1)
}
