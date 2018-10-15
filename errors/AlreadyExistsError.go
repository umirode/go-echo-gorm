package errors

import (
	"net/http"
)

type AlreadyExistsError struct {
	Status int
}

func NewAlreadyExistsError() *AlreadyExistsError {
	return &AlreadyExistsError{
		Status: http.StatusConflict,
	}
}

func (e *AlreadyExistsError) Error() string {
	return "Already exists"
}
