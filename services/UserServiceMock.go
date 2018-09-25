package services

import (
    "github.com/umirode/go-rest/models"
    "github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
    mock.Mock

    IUserService
}

func (s *UserServiceMock) GetUserByID(id uint) *models.UserModel {
    args := s.Called(id)

    return args.Get(0).(*models.UserModel)
}

func (s *UserServiceMock) GetAllUsers() *[]models.UserModel {
    args := s.Called()

    return args.Get(0).(*[]models.UserModel)
}

func (s *UserServiceMock) CreateUser(user *models.UserModel) error {
    args := s.Called(user)

    return args.Error(0)
}

func (s *UserServiceMock) UpdateUser(id uint, user *models.UserModel) error {
    args := s.Called(id, user)

    return args.Error(0)
}

func (s *UserServiceMock) DeleteUser(id uint) error {
    args := s.Called(id)

    return args.Error(0)
}
