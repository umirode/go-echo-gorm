package services

import (
    "github.com/umirode/go-rest/repositories"
    "github.com/umirode/go-rest/models"
    "github.com/umirode/go-rest/specifications"
)

type UserService struct {
    IUserService

    BaseService

    Repository repositories.IUserRepository
}

func (s *UserService) GetUserByID(id uint) *models.UserModel {
    idSpecification := &specifications.IdSpecification{
        Id: id,
    }

    users := s.Repository.Query(idSpecification)
    if len(*users) < 1 {
        return nil
    }

    return &(*users)[0]
}

func (s *UserService) GetAllUsers() *[]models.UserModel {
    users := s.Repository.Query()

    return users
}

func (s *UserService) CreateUser(user *models.UserModel) error {
    err := s.Repository.AddUser(user)

    return err
}

func (s *UserService) UpdateUser(id uint, user *models.UserModel) error {
    user.ID = id

    err := s.Repository.UpdateUser(user)

    return err
}

func (s *UserService) DeleteUser(id uint) error {
    user := new(models.UserModel)
    user.ID = id

    err := s.Repository.DeleteUser(user)

    return err
}
