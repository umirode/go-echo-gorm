package errors

import (
	"net/http"
)

type AuthError struct{}

func NewAuthError() *AuthError {
	return &AuthError{}
}

func (e *AuthError) Status() int {
	return http.StatusUnauthorized
}

func (e *AuthError) Error() string {
	return "User unauthorized"
}
