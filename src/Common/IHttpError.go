package Common

type IHttpError interface {
	error
	Status() int
}
