package Error

import (
	"net/http"
)

type InvalidError struct {
}

func (*InvalidError) Error() string {
	return "Invalid"
}

func (*InvalidError) Status() int {
	return http.StatusBadRequest
}

func NewInvalidError() *InvalidError {
	return &InvalidError{}
}
