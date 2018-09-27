package repositories

import (
    "errors"
    "github.com/jinzhu/gorm"
)

type BaseDatabaseRepository struct {
    Database *gorm.DB
}

func (r *BaseDatabaseRepository) create(i interface{}) error {
    if !r.Database.NewRecord(i) {
        return r.getAlreadyExistsError()
    }

    r.Database.Create(i)

    return nil
}

func (r *BaseDatabaseRepository) update(i interface{}, fields ...string) error {
    if r.Database.NewRecord(i) {
        return r.getNotExistsError()
    }

    result := r.Database.Model(i).Select(fields).Updates(i)
    if result.RecordNotFound() {
        return r.getNotExistsError()
    }

    return nil
}

func (r *BaseDatabaseRepository) delete(i interface{}) error {
    if r.Database.NewRecord(i) {
        return r.getNotExistsError()
    }

    result := r.Database.Delete(i)
    if result.RecordNotFound() {
        return r.getNotExistsError()
    }

    return nil
}

func (r *BaseDatabaseRepository) getNotExistsError() error {
    return errors.New("Not exists ")
}

func (r *BaseDatabaseRepository) getAlreadyExistsError() error {
    return errors.New("Already exists ")
}
