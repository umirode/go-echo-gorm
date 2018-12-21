package Converter

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Infrastructure/Persistence/Gorm/Model"
)

type UserConverter struct {
}

func (*UserConverter) ToDatabaseEntity(entity interface{}) (interface{}, error) {
	if entity == nil {
		return nil, nil
	}

	user := entity.(*Entity.User)

	model := &Model.User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}

	birthdayConverter := &BirthdayConverter{}

	for _, birthday := range user.Birthdays {
		birthdayModel, _ := birthdayConverter.ToDatabaseEntity(birthday)

		if birthdayModel != nil {
			model.Birthdays = append(model.Birthdays, birthdayModel.(*Model.Birthday))
		}
	}

	return model, nil
}

func (*UserConverter) ToAppEntity(entity interface{}) (interface{}, error) {
	if entity == nil {
		return nil, nil
	}

	user := entity.(*Model.User)

	model := &Entity.User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}

	birthdayConverter := &BirthdayConverter{}

	for _, birthday := range user.Birthdays {
		birthdayEntity, _ := birthdayConverter.ToAppEntity(birthday)

		if birthdayEntity != nil {
			model.Birthdays = append(model.Birthdays, birthdayEntity.(*Entity.Birthday))
		}
	}

	return model, nil
}
