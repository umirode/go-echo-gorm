package repositories

import (
	"github.com/umirode/go-rest/database"
	"github.com/umirode/go-rest/models"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) GetAllUsers() []models.UserModel {
	var users []models.UserModel
	database.GetDBConnection().Find(&users)

	return users
}

func (ur *UserRepository) GetSingleUser(id uint) models.UserModel {
	user := models.UserModel{
		ID: id,
	}
	database.GetDBConnection().First(&user)

	return user
}

func (ur *UserRepository) CreateUser(name string) models.UserModel {
	user := models.UserModel{
		Name: name,
	}
	database.GetDBConnection().Create(&user)

	return user
}

func (ur *UserRepository) UpdateUser(id uint, name string) models.UserModel {
	user := models.UserModel{
		ID: id,
	}
	database.GetDBConnection().First(&user)

	user.Name = name
	database.GetDBConnection().Save(&user)

	return user
}

func (ur *UserRepository) DeleteUser(id uint) {
	user := models.UserModel{
		ID: id,
	}
	database.GetDBConnection().Delete(&user)
}
