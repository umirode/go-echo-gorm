package generator

import (
	"os"
	"path"
)

type FileCreator struct{}

func (fc *FileCreator) Create(filePath string) (*os.File, error) {
	directory, _ := path.Split(filePath)
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return nil, err
	}

	newFile, err := os.Create(filePath)
	if err != nil {
		os.Remove(filePath)
		return nil, err
	}

	return newFile, nil
}
