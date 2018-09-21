package services

import (
	"github.com/umirode/go-rest/models"
)

type IUserService interface {
	GetUsersByID(id uint) *[]models.UserModel
	GetUsersByName(name string) *[]models.UserModel
	GetAllUsers() *[]models.UserModel

	CreateUser(user *models.UserModel) error
	UpdateUser(id uint, user *models.UserModel) error
	DeleteUser(id uint) error
}
