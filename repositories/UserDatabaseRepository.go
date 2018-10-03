package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/go-rest/models"
)

type UserDatabaseRepository struct {
	BaseDatabaseRepository

	IUserRepository
}

func NewUserDatabaseRepository(database *gorm.DB) *UserDatabaseRepository {
	repository := &UserDatabaseRepository{}

	repository.Database = database

	return repository
}

func (r *UserDatabaseRepository) FindAll() *[]models.UserModel {
	users := make([]models.UserModel, 0)

	r.Database.Find(&users)

	return &users
}

func (r *UserDatabaseRepository) FindSingleById(id uint) *models.UserModel {
	user := new(models.UserModel)

	r.Database.Where("id = ?", id).First(&user)

	return user
}

func (r *UserDatabaseRepository) AddUser(user *models.UserModel) error {
	err := r.create(user)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserDatabaseRepository) UpdateUser(user *models.UserModel) error {
	err := r.update(user, map[string]interface{}{
		"FIELD": "TEST",
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *UserDatabaseRepository) DeleteUser(user *models.UserModel) error {
	err := r.delete(user)
	if err != nil {
		return err
	}

	return nil
}
