package repositories

import (
	"github.com/stretchr/testify/mock"
	"github.com/umirode/go-rest/models"
)

type MockUserRepository struct {
	mock.Mock

	IUserRepository
}

func (r *MockUserRepository) FindAll() *[]models.UserModel {
	args := r.Called()

	return args.Get(0).(*[]models.UserModel)
}

func (r *MockUserRepository) FindSingleById(id uint) (*models.UserModel, error) {
	args := r.Called(id)

	return args.Get(0).(*models.UserModel), args.Error(1)
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
