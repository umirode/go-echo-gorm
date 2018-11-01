package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/umirode/go-rest/models"
	"github.com/umirode/go-rest/repositories"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	userRepository := &repositories.MockUserRepository{}
	userRepository.On("AddUser", &models.UserModel{}).Return(nil)

	s := &UserService{
		UserRepository: userRepository,
	}

	err := s.CreateUser(&models.UserModel{})

	assert.NoError(t, err)
}

func TestUserService_DeleteUser(t *testing.T) {
	userRepository := &repositories.MockUserRepository{}
	userRepository.On("DeleteUser", &models.UserModel{}).Return(nil)

	s := &UserService{
		UserRepository: userRepository,
	}

	err := s.DeleteUser(0)

	assert.NoError(t, err)
}

func TestUserService_GetAllUsers(t *testing.T) {
	userRepository := &repositories.MockUserRepository{}
	userRepository.On("FindAll").Return(&[]models.UserModel{{}})

	s := &UserService{
		UserRepository: userRepository,
	}

	users, err := s.GetAllUsers()

	assert.NotEmpty(t, users)
	assert.NoError(t, err)
}

func TestUserService_GetUserByID(t *testing.T) {
	userRepository := &repositories.MockUserRepository{}
	userRepository.On("FindSingleByID", mock.AnythingOfType("uint")).Return(&models.UserModel{
		ID: 1,
	}, nil)

	s := &UserService{
		UserRepository: userRepository,
	}

	user, err := s.GetUserByID(1)

	assert.NotEmpty(t, user)
	assert.NoError(t, err)
}

func TestUserService_UpdateUser(t *testing.T) {
	userRepository := &repositories.MockUserRepository{}
	userRepository.On("UpdateUser", &models.UserModel{}, mock.Anything).Return(nil)

	s := &UserService{
		UserRepository: userRepository,
	}

	err := s.UpdateUser(0, &models.UserModel{})

	assert.NoError(t, err)
}
