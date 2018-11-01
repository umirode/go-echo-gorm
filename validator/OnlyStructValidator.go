package validator

type OnlyStructValidator struct {
}

func NewOnlyStructValidator() *OnlyStructValidator {
	return &OnlyStructValidator{}
}

func (OnlyStructValidator) Validate(i interface{}) error {
	v := NewValidator()

	return v.Validator.Struct(i)
}
