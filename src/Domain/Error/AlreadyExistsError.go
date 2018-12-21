package Error

import (
	"net/http"
)

type AlreadyExistsError struct {
}

func (*AlreadyExistsError) Error() string {
	return "Already exists"
}

func (*AlreadyExistsError) Status() int {
	return http.StatusConflict
}

func NewAlreadyExistsError() *AlreadyExistsError {
	return &AlreadyExistsError{}
}
