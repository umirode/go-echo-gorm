package Repository

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Infrastructure/Persistence/Gorm/Converter"
	"github.com/umirode/go-rest/src/Infrastructure/Persistence/Gorm/Model"
)

type BirthdayRepository struct {
	db                *gorm.DB
	birthdayConverter *Converter.BirthdayConverter
	userConverter     *Converter.UserConverter
}

func NewBirthdayRepository(db *gorm.DB) *BirthdayRepository {
	return &BirthdayRepository{
		birthdayConverter: &Converter.BirthdayConverter{},
		userConverter:     &Converter.UserConverter{},
		db:                db,
	}
}

func (r *BirthdayRepository) Save(birthday *Entity.Birthday) error {
	model, _ := r.birthdayConverter.ToDatabaseEntity(birthday)
	if model == nil {
		return nil
	}

	r.db.Save(model)

	return nil
}

func (r *BirthdayRepository) Delete(birthday *Entity.Birthday) error {
	model, _ := r.birthdayConverter.ToDatabaseEntity(birthday)
	if model == nil {
		return nil
	}

	r.db.Delete(model)

	return nil
}

func (r *BirthdayRepository) FindAllByUser(user *Entity.User) ([]*Entity.Birthday, error) {
	model, _ := r.userConverter.ToDatabaseEntity(user)
	if model == nil {
		return make([]*Entity.Birthday, 0), nil
	}

	r.db.Model(model).Related(&model.(*Model.User).Birthdays)

	entity, _ := r.userConverter.ToAppEntity(model)
	if entity == nil {
		return make([]*Entity.Birthday, 0), nil
	}

	return entity.(*Entity.User).Birthdays, nil
}

func (r *BirthdayRepository) FindOneById(id uint) (*Entity.Birthday, error) {
	model := &Model.Birthday{}

	r.db.Where("id = ?", id).First(model)
	if r.db.NewRecord(model) {
		return nil, nil
	}

	entity, _ := r.birthdayConverter.ToAppEntity(model)
	if entity == nil {
		return nil, nil
	}

	return entity.(*Entity.Birthday), nil
}
