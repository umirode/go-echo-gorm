package repositories

import (
    "github.com/umirode/go-rest/models"
    "github.com/stretchr/testify/mock"
    "github.com/umirode/go-rest/specifications"
    "github.com/jinzhu/gorm"
)

type MockUserRepository struct {
    mock.Mock

    IUserRepository
}

func (r *MockUserRepository) Query(specifications ...specifications.IDatabaseSpecification) *[]models.UserModel {
    args := r.Called()

    return args.Get(0).(*[]models.UserModel)
}

func (r *MockUserRepository) RawQuery(handler func(db *gorm.DB, users *[]models.UserModel)) *[]models.UserModel {
    args := r.Called(handler)

    return args.Get(0).(*[]models.UserModel)
}

func (r *MockUserRepository) AddUser(user *models.UserModel) error {
    args := r.Called(user)

    return args.Error(0)
}

func (r *MockUserRepository) UpdateUser(user *models.UserModel) error {
    args := r.Called(user)

    return args.Error(1)
}

func (r *MockUserRepository) DeleteUser(user *models.UserModel) error {
    args := r.Called(user)

    return args.Error(0)
}
