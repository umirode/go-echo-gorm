package router

import (
	"gopkg.in/go-playground/validator.v9"
)

type structValidator struct{}

func (structValidator) Validate(i interface{}) error {
	return validator.New().Struct(i)
}
