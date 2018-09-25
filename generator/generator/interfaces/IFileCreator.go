package interfaces

import (
    "os"
)

type IFileCreator interface {
    Create(filePath string) (*os.File, error)
}
