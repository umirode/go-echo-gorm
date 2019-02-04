package Hydrator

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

type BirthdayHydrator struct {
}

func (*BirthdayHydrator) Create(data map[string]interface{}) (interface{}, error) {
	panic("implement me")
}

func (*BirthdayHydrator) Hydrate(data map[string]interface{}, object interface{}) (interface{}, error) {
	panic("implement me")
}

func (*BirthdayHydrator) Extract(object interface{}) (map[string]interface{}, error) {
	birthday := object.(*Entity.Birthday)

	return map[string]interface{}{
		"id":     birthday.ID,
		"month":  birthday.Month,
		"number": birthday.Number,
		"name":   birthday.Name,
	}, nil
}
