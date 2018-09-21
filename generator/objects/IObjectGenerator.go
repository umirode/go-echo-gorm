package objects

type IObjectGenerator interface {
	Generate(name string, args []string) error
}
