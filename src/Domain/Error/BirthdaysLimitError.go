package Error

import (
	"net/http"
)

type BirthdaysLimitError struct {
}

func (*BirthdaysLimitError) Error() string {
	return "Type convert error"
}

func (*BirthdaysLimitError) Status() int {
	return http.StatusBadRequest
}

func NewBirthdaysLimitError() *BirthdaysLimitError {
	return &BirthdaysLimitError{}
}
