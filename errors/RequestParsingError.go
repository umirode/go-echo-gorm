package errors

import (
	"net/http"
)

type RequestParsingError struct {
	Status int
}

func NewRequestParsingError() *RequestParsingError {
	return &RequestParsingError{
		Status: http.StatusBadRequest,
	}
}

func (e *RequestParsingError) Error() string {
	return "Parsing request error"
}
