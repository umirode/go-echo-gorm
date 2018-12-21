package Repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

type BirthdayRepositoryMock struct {
	mock.Mock
}

func (r *BirthdayRepositoryMock) Save(birthday *Entity.Birthday) error {
	args := r.Called(birthday)

	return args.Error(0)
}

func (r *BirthdayRepositoryMock) Delete(birthday *Entity.Birthday) error {
	args := r.Called(birthday)

	return args.Error(0)
}

func (r *BirthdayRepositoryMock) FindAllByUser(user *Entity.User) ([]*Entity.Birthday, error) {
	args := r.Called(user)

	birthdays, ok := args.Get(0).([]*Entity.Birthday)
	if ok {
		return birthdays, args.Error(1)
	}

	return nil, args.Error(1)
}

func (r *BirthdayRepositoryMock) FindOneById(id uint) (*Entity.Birthday, error) {
	args := r.Called(id)

	birthday, ok := args.Get(0).(*Entity.Birthday)
	if ok {
		return birthday, args.Error(1)
	}

	return nil, args.Error(1)
}
