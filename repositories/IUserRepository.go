package repositories

import (
	"github.com/umirode/go-rest/models"
	"github.com/umirode/go-rest/specifications"
)

type IUserRepository interface {
	AddUser(user *models.UserModel) error
	UpdateUser(user *models.UserModel) error
	DeleteUser(user *models.UserModel) error

	Query(specifications ...specifications.IDatabaseSpecification) *[]models.UserModel
}
