package Repository

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Infrastructure/Persistence/Gorm/Converter"
	"github.com/umirode/go-rest/src/Infrastructure/Persistence/Gorm/Model"
)

type UserRepository struct {
	db            *gorm.DB
	userConverter *Converter.UserConverter
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		userConverter: &Converter.UserConverter{},
		db:            db,
	}
}

func (r *UserRepository) Save(user *Entity.User) error {
	model, _ := r.userConverter.ToDatabaseEntity(user)
	if model == nil {
		return nil
	}

	r.db.Save(model)

	return nil
}

func (r *UserRepository) Delete(user *Entity.User) error {
	model, _ := r.userConverter.ToDatabaseEntity(user)
	if model == nil {
		return nil
	}

	r.db.Delete(model)

	return nil
}

func (r *UserRepository) FindAll() ([]*Entity.User, error) {
	models := make([]*Model.User, 0)

	r.db.Find(&models)

	entities := make([]*Entity.User, 0)
	for _, model := range models {
		r.db.Model(model).Related(&model.Birthdays)

		entity, _ := r.userConverter.ToAppEntity(model)

		entities = append(entities, entity.(*Entity.User))
	}

	return entities, nil
}

func (r *UserRepository) FindOneByID(id uint) (*Entity.User, error) {
	model := &Model.User{}

	r.db.Where("id = ?", id).First(model)
	if r.db.NewRecord(model) {
		return nil, nil
	}

	r.db.Model(model).Related(&model.Birthdays)

	entity, _ := r.userConverter.ToAppEntity(model)
	if entity == nil {
		return nil, nil
	}

	return entity.(*Entity.User), nil
}

func (r *UserRepository) FindOneByEmail(email string) (*Entity.User, error) {
	model := &Model.User{}

	r.db.Where("email = ?", email).First(model)
	if r.db.NewRecord(model) {
		return nil, nil
	}

	r.db.Model(model).Related(&model.Birthdays)

	entity, _ := r.userConverter.ToAppEntity(model)
	if entity == nil {
		return nil, nil
	}

	return entity.(*Entity.User), nil
}

func (r *UserRepository) FindOneByEmailAndPassword(email string, password string) (*Entity.User, error) {
	model := &Model.User{}

	r.db.Where("email = ? and password_hash = ?", email, password).First(model)
	if r.db.NewRecord(model) {
		return nil, nil
	}

	r.db.Model(model).Related(&model.Birthdays)

	entity, _ := r.userConverter.ToAppEntity(model)
	if entity == nil {
		return nil, nil
	}

	return entity.(*Entity.User), nil
}
