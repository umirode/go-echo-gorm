package errors

import (
	"net/http"
)

type PasswordDoesNotMatchConfirmPasswordError struct{}

func NewPasswordDoesNotMatchConfirmPasswordError() *PasswordDoesNotMatchConfirmPasswordError {
	return &PasswordDoesNotMatchConfirmPasswordError{}
}

func (e *PasswordDoesNotMatchConfirmPasswordError) Status() int {
	return http.StatusUnauthorized
}

func (e *PasswordDoesNotMatchConfirmPasswordError) Error() string {
	return "Password does not match the confirm password"
}
