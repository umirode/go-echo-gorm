package errors

import (
	"net/http"
)

type NotFoundError struct{}

func NewNotFoundError() *NotFoundError {
	return &NotFoundError{}
}

func (e *NotFoundError) Status() int {
	return http.StatusNotFound
}

func (e *NotFoundError) Error() string {
	return "Not found"
}
