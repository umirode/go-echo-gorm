package Error

import (
	"net/http"
)

type AccessError struct {
}

func (*AccessError) Error() string {
	return "Not found"
}

func (*AccessError) Status() int {
	return http.StatusForbidden
}

func NewAccessError() *AccessError {
	return &AccessError{}
}
