package Repository

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(user *Entity.User) error {
	r.db.Save(user)

	return nil
}

func (r *UserRepository) FindOneByID(id uint) (*Entity.User, error) {
	user := &Entity.User{}

	r.db.Where("id = ?", id).First(user)
	if r.db.NewRecord(user) {
		return nil, nil
	}

	return user, nil
}

func (r *UserRepository) FindOneByEmail(email string) (*Entity.User, error) {
	user := &Entity.User{}

	r.db.Where("email = ?", email).First(user)
	if r.db.NewRecord(user) {
		return nil, nil
	}

	return user, nil
}

func (r *UserRepository) FindOneByEmailAndPassword(email string, password string) (*Entity.User, error) {
	user := &Entity.User{}

	r.db.Where("email = ? and password_hash = ?", email, password).First(user)
	if r.db.NewRecord(user) {
		return nil, nil
	}

	return user, nil
}
