package errors

import (
	"net/http"
)

type AlreadyExistsError struct{}

func NewAlreadyExistsError() *AlreadyExistsError {
	return &AlreadyExistsError{}
}

func (e *AlreadyExistsError) Status() int {
	return http.StatusConflict
}

func (e *AlreadyExistsError) Error() string {
	return "Already exists"
}
