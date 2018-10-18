package errors

import (
	"net/http"
)

type InvalidOldPasswordError struct{}

func NewInvalidOldPasswordError() *InvalidOldPasswordError {
	return &InvalidOldPasswordError{}
}

func (e *InvalidOldPasswordError) Status() int {
	return http.StatusUnauthorized
}

func (e *InvalidOldPasswordError) Error() string {
	return "Invalid old password"
}
