package errors

type IHTTPError interface {
	Error() string
	Status() int
}
