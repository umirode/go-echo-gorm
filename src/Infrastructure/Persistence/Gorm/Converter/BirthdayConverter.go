package Converter

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Infrastructure/Persistence/Gorm/Model"
	"time"
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
		Date:   birthday.Date.String(),
		UserID: birthday.User.ID,
	}

	return model, nil
}

func (*BirthdayConverter) ToAppEntity(entity interface{}) (interface{}, error) {
	if entity == nil {
		return nil, nil

	}

	birthday := entity.(*Model.Birthday)

	date, _ := time.Parse(birthday.Date, "2006-01-02 15:04:05.999999999 -0700 MST")

	model := &Entity.Birthday{
		ID:   birthday.ID,
		Name: birthday.Name,
		Date: date,
		User: &Entity.User{
			ID: birthday.ID,
		},
	}

	return model, nil
}
