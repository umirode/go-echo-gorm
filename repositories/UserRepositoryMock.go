package repositories

import (
	"github.com/stretchr/testify/mock"
	"github.com/umirode/go-rest/models"
)

type MockUserRepository struct {
	mock.Mock
}

func (r *MockUserRepository) FindSingleByEmailAndPassword(email string, password string) (*models.UserModel, error) {
	args := r.Called(email, password)

	user, ok := args.Get(0).(*models.UserModel)
	if ok {
		return user, args.Error(1)
	}

	return nil, args.Error(1)
}

func (r *MockUserRepository) FindAll() *[]models.UserModel {
	args := r.Called()

	users, ok := args.Get(0).(*[]models.UserModel)
	if ok {
		return users
	}

	return nil
}

func (r *MockUserRepository) FindSingleByID(id uint) (*models.UserModel, error) {
	args := r.Called(id)

	user, ok := args.Get(0).(*models.UserModel)
	if ok {
		return user, args.Error(1)
	}

	return nil, args.Error(1)
}

func (r *MockUserRepository) AddUser(user *models.UserModel) error {
	args := r.Called(user)

	return args.Error(0)
}

func (r *MockUserRepository) UpdateUser(user *models.UserModel) error {
	args := r.Called(user)

	return args.Error(0)
}

func (r *MockUserRepository) DeleteUser(user *models.UserModel) error {
	args := r.Called(user)

	return args.Error(0)
}
