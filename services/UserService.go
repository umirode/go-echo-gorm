package services

import (
	"github.com/umirode/go-rest/models"
	"github.com/umirode/go-rest/repositories"
)

type UserService struct {
	BaseService

	UserRepository repositories.IUserRepository
}

func (s *UserService) GetUserByID(id uint) (*models.UserModel, error) {
	user, err := s.UserRepository.FindSingleByID(id)

	return user, err
}

func (s *UserService) GetAllUsers() (*[]models.UserModel, error) {
	users := s.UserRepository.FindAll()

	return users, nil
}

func (s *UserService) CreateUser(user *models.UserModel) error {
	err := s.UserRepository.AddUser(user)

	return err
}

func (s *UserService) UpdateUser(id uint, user *models.UserModel) error {
	user.ID = id

	err := s.UserRepository.UpdateUser(user)

	return err
}

func (s *UserService) DeleteUser(id uint) error {
	user := new(models.UserModel)
	user.ID = id

	err := s.UserRepository.DeleteUser(user)

	return err
}
