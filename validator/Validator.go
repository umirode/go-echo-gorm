package validator

import (
	"github.com/umirode/go-rest/validator/validators"
	"gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	Validator *validator.Validate
}

func NewValidator() *Validator {
	v := &Validator{
		Validator: validator.New(),
	}

	v.registerCustomValidator(validators.NewDateCustomValidator())

	return v
}

func (v *Validator) registerCustomValidator(customValidator validators.ICustomValidator) {
	v.Validator.RegisterValidation(customValidator.GetTag(), customValidator.GetValidator)
}
