package Validator

import (
	"github.com/sirupsen/logrus"
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
	err := v.Validator.RegisterValidation(customValidator.GetTag(), customValidator.GetValidator)
	if err != nil {
		logrus.Error(err)
	}
}
