package Error

import (
	"net/http"
)

type TypeError struct {
}

func (*TypeError) Error() string {
	return "Type convert error"
}

func (*TypeError) Status() int {
	return http.StatusInternalServerError
}

func NewTypeError() *TypeError {
	return &TypeError{}
}
