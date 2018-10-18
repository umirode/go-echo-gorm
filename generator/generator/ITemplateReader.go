package generator

type ITemplateReader interface {
	Read(filePath string) ([]byte, error)
}
