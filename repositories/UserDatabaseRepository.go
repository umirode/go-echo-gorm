package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/go-rest/errors"
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

func (r *UserDatabaseRepository) FindSingleById(id uint) (*models.UserModel, error) {
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

func (r *UserDatabaseRepository) UpdateUser(user *models.UserModel) error {
	err := r.update(user, map[string]interface{}{
		"name": user.Name,
	})

	return err
}

func (r *UserDatabaseRepository) DeleteUser(user *models.UserModel) error {
	err := r.delete(user)

	return err
}
