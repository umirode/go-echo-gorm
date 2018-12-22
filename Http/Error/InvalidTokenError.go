package Error

import (
	"net/http"
)

type InvalidTokenError struct{}

func NewInvalidTokenError() *InvalidTokenError {
	return &InvalidTokenError{}
}

func (e *InvalidTokenError) Status() int {
	return http.StatusUnauthorized
}

func (e *InvalidTokenError) Error() string {
	return "Invalid or expired jwt"
}
