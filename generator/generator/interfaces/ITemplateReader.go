package interfaces

type ITemplateReader interface {
    Read(filePath string) ([]byte, error)
}
