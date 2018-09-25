package generator

import (
    "io/ioutil"
)

type TemplateReader struct{}

func (tr *TemplateReader) Read(filePath string) ([]byte, error) {
    return ioutil.ReadFile(filePath)
}
