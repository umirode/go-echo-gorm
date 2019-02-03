package Error

import (
	"net/http"
)

type AccessError struct {
}

func (*AccessError) Error() string {
	return "Access error"
}

func (*AccessError) Status() int {
	return http.StatusForbidden
}

func NewAccessError() *AccessError {
	return &AccessError{}
}
