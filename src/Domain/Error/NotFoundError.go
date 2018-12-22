package Error

import (
	"net/http"
)

type NotFoundError struct {
}

func (*NotFoundError) Error() string {
	return "Not found"
}

func (*NotFoundError) Status() int {
	return http.StatusNotFound
}

func NewNotFoundError() *NotFoundError {
	return &NotFoundError{}
}
