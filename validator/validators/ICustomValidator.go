package validators

import (
	"gopkg.in/go-playground/validator.v9"
)

type ICustomValidator interface {
	GetValidator(fl validator.FieldLevel) bool
	GetTag() string
}
