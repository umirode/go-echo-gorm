package generator

import (
	"os"
)

type FileCreator struct{}

func (fc *FileCreator) Create(filePath string) (*os.File, error) {
	newFile, err := os.Create(filePath)
	if err != nil {
		os.Remove(filePath)
		return nil, err
	}

	return newFile, nil
}
