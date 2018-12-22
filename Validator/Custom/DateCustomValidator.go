package Custom

import (
	"gopkg.in/go-playground/validator.v9"
	"time"
)

type DateCustomValidator struct{}

func NewDateCustomValidator() *DateCustomValidator {
	return &DateCustomValidator{}
}

func (v *DateCustomValidator) GetValidator(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC822, fl.Field().String())
	if err != nil {
		return false
	}

	return false
}

func (v *DateCustomValidator) GetTag() string {
	return "date"
}
