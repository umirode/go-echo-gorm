package Repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (r *UserRepositoryMock) Save(user *Entity.User) error {
	args := r.Called(user)

	return args.Error(0)
}

func (r *UserRepositoryMock) Delete(user *Entity.User) error {
	args := r.Called(user)

	return args.Error(0)
}

func (r *UserRepositoryMock) FindAll() ([]*Entity.User, error) {
	args := r.Called()

	users, ok := args.Get(0).([]*Entity.User)
	if ok {
		return users, args.Error(1)
	}

	return nil, args.Error(1)
}

func (r *UserRepositoryMock) FindOneByID(id uint) (*Entity.User, error) {
	args := r.Called(id)

	user, ok := args.Get(0).(*Entity.User)
	if ok {
		return user, args.Error(1)
	}

	return nil, args.Error(1)
}

func (r *UserRepositoryMock) FindOneByEmail(email string) (*Entity.User, error) {
	args := r.Called(email)

	user, ok := args.Get(0).(*Entity.User)
	if ok {
		return user, args.Error(1)
	}

	return nil, args.Error(1)
}

func (r *UserRepositoryMock) FindOneByEmailAndPassword(email string, password string) (*Entity.User, error) {
	args := r.Called(email, password)

	user, ok := args.Get(0).(*Entity.User)
	if ok {
		return user, args.Error(1)
	}

	return nil, args.Error(1)
}
