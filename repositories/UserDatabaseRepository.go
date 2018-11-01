package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/go-rest/errors"
	"github.com/umirode/go-rest/models"
)

type UserDatabaseRepository struct {
	BaseDatabaseRepository
}

func NewUserDatabaseRepository(database *gorm.DB) *UserDatabaseRepository {
	repository := &UserDatabaseRepository{}

	repository.Database = database

	return repository
}

func (r *UserDatabaseRepository) FindSingleByEmail(email string) (*models.UserModel, error) {
	user := new(models.UserModel)

	r.Database.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		return nil, errors.NewNotFoundError()
	}

	return user, nil
}

func (r *UserDatabaseRepository) FindSingleByEmailAndPassword(email string, password string) (*models.UserModel, error) {
	user := new(models.UserModel)

	r.Database.Where("email = ? AND password = ?", email, password).First(&user)

	if user.ID == 0 {
		return nil, errors.NewNotFoundError()
	}

	return user, nil
}

func (r *UserDatabaseRepository) FindAll() *[]models.UserModel {
	users := make([]models.UserModel, 0)

	r.Database.Find(&users)

	return &users
}

func (r *UserDatabaseRepository) FindSingleByID(id uint) (*models.UserModel, error) {
	user := new(models.UserModel)

	r.Database.Where("id = ?", id).First(&user)

	if user.ID == 0 {
		return nil, errors.NewNotFoundError()
	}

	return user, nil
}

func (r *UserDatabaseRepository) AddUser(user *models.UserModel) error {
	err := r.create(user)

	return err
}

func (r *UserDatabaseRepository) UpdateUser(user *models.UserModel, data map[string]interface{}) error {
	err := r.update(user, data)

	return err
}

func (r *UserDatabaseRepository) DeleteUserByEmail(email string) error {
	result := r.Database.Where("email = ?", email).Delete(models.UserModel{})

	return result.Error
}

func (r *UserDatabaseRepository) DeleteUser(user *models.UserModel) error {
	err := r.delete(user)

	return err
}
