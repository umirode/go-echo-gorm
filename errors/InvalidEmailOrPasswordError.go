package errors

import (
	"net/http"
)

type InvalidEmailOrPasswordError struct{}

func NewInvalidEmailOrPasswordError() *InvalidEmailOrPasswordError {
	return &InvalidEmailOrPasswordError{}
}

func (e *InvalidEmailOrPasswordError) Status() int {
	return http.StatusUnauthorized
}

func (e *InvalidEmailOrPasswordError) Error() string {
	return "Invalid email or password"
}
