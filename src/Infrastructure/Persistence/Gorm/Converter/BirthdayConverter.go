package Converter

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Infrastructure/Persistence/Gorm/Model"
)

type BirthdayConverter struct {
}

func (*BirthdayConverter) ToDatabaseEntity(entity interface{}) (interface{}, error) {
	if entity == nil {
		return nil, nil
	}

	birthday := entity.(*Entity.Birthday)

	model := &Model.Birthday{
		ID:     birthday.ID,
		Name:   birthday.Name,
		Date:   birthday.Date,
		UserID: birthday.User.ID,
	}

	return model, nil
}

func (*BirthdayConverter) ToAppEntity(entity interface{}) (interface{}, error) {
	if entity == nil {
		return nil, nil

	}

	birthday := entity.(*Model.Birthday)

	model := &Entity.Birthday{
		ID:   birthday.ID,
		Name: birthday.Name,
		Date: birthday.Date,
		User: &Entity.User{
			ID: birthday.ID,
		},
	}

	return model, nil
}
