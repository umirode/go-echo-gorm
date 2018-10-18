package router

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type structForValidate struct {
	Test string `validate:"required"`
}

func TestStructValidator_Validate(t *testing.T) {
	validator := &structValidator{}

	err := validator.Validate(&structForValidate{
		Test: "test",
	})

	assert.NoError(t, err)
}

func TestStructValidator_Validate_Error(t *testing.T) {
	validator := &structValidator{}

	err := validator.Validate(&structForValidate{})

	assert.Error(t, err)
}
