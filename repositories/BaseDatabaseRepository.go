package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/go-rest/errors"
)

type BaseDatabaseRepository struct {
	Database *gorm.DB
}

func (r *BaseDatabaseRepository) create(i interface{}) error {
	if !r.Database.NewRecord(i) {
		return errors.NewAlreadyExistsError()
	}

	result := r.Database.Create(i)

	return result.Error
}

func (r *BaseDatabaseRepository) update(i interface{}, data interface{}) error {
	if r.Database.NewRecord(i) {
		return errors.NewNotFoundError()
	}

	result := r.Database.Model(i).Updates(data)
	if result.RowsAffected == 0 {
		return errors.NewNotFoundError()
	}

	return result.Error
}

func (r *BaseDatabaseRepository) delete(i interface{}) error {
	if r.Database.NewRecord(i) {
		return errors.NewNotFoundError()
	}

	result := r.Database.Delete(i)
	if result.RowsAffected == 0 {
		return errors.NewNotFoundError()
	}

	return result.Error
}
