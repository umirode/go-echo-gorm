package validators

import (
	"gopkg.in/go-playground/validator.v9"
	"regexp"
)

type DateCustomValidator struct{}

func NewDateCustomValidator() *DateCustomValidator {
	return &DateCustomValidator{}
}

func (v *DateCustomValidator) GetValidator(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`(?m)^\s*(3[01]|[12][0-9]|0?[1-9])\.(1[012]|0?[1-9])\.((?:19|20)\d{2})`)

	result := re.MatchString(fl.Field().String())
	if result {
		return true
	}

	return false
}

func (v *DateCustomValidator) GetTag() string {
	return "date"
}
