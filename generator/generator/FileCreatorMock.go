package generator

import (
	"github.com/stretchr/testify/mock"
	"os"
)

type FileCreatorMock struct {
	mock.Mock
}

func (fc *FileCreatorMock) Create(filePath string) (*os.File, error) {
	args := fc.Called(filePath)

	return args.Get(0).(*os.File), args.Error(1)
}
