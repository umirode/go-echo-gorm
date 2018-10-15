package errors

import (
	"net/http"
)

type NotFoundError struct {
	Status int
}

func NewNotFoundError() *NotFoundError {
	return &NotFoundError{
		Status: http.StatusNotFound,
	}
}

func (e *NotFoundError) Error() string {
	return "Not found"
}
