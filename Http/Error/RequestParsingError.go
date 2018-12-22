package Error

import (
	"net/http"
)

type RequestParsingError struct{}

func NewRequestParsingError() *RequestParsingError {
	return &RequestParsingError{}
}

func (e *RequestParsingError) Status() int {
	return http.StatusBadRequest
}

func (e *RequestParsingError) Error() string {
	return "Parsing request error"
}
