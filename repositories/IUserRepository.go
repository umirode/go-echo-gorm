package repositories

import (
	"github.com/umirode/go-rest/models"
)

type IUserRepository interface {
	AddUser(user *models.UserModel) error
	UpdateUser(user *models.UserModel) error
	DeleteUser(user *models.UserModel) error
	DeleteUserByEmail(email string) error

	FindAll() *[]models.UserModel
	FindSingleByID(id uint) (*models.UserModel, error)
	FindSingleByEmailAndPassword(email string, password string) (*models.UserModel, error)
	FindSingleByEmail(email string) (*models.UserModel, error)
}
