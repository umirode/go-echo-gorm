package Validator

import (
	"github.com/umirode/go-rest/Validator/Custom"
	"gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	Validator *validator.Validate
}

func NewValidator() *Validator {
	v := &Validator{
		Validator: validator.New(),
	}

	return v
}

func (v *Validator) registerCustomValidator(customValidator Custom.ICustomValidator) {
	v.Validator.RegisterValidation(customValidator.GetTag(), customValidator.GetValidator)
}
