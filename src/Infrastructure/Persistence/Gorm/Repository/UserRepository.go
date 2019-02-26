package Repository

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

type UserRepository struct {
	BaseRepository
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Save(user *Entity.User) error {
	r.GetGormDB().Save(user)

	return nil
}

func (r *UserRepository) FindOneByID(id uint) (*Entity.User, error) {
	user := &Entity.User{}

	r.GetGormDB().Where("id = ?", id).First(user)
	if r.GetGormDB().NewRecord(user) {
		return nil, nil
	}

	return user, nil
}

func (r *UserRepository) FindOneByEmail(email string) (*Entity.User, error) {
	user := &Entity.User{}

	r.GetGormDB().Where("email = ?", email).First(user)
	if r.GetGormDB().NewRecord(user) {
		return nil, nil
	}

	return user, nil
}

func (r *UserRepository) FindOneByEmailAndPassword(email string, password string) (*Entity.User, error) {
	user := &Entity.User{}

	r.GetGormDB().Where("email = ? and password_hash = ?", email, password).First(user)
	if r.GetGormDB().NewRecord(user) {
		return nil, nil
	}

	return user, nil
}
