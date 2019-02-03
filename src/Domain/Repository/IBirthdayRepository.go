package Repository

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

type IBirthdayRepository interface {
	Save(birthday *Entity.Birthday) error
	Delete(birthday *Entity.Birthday) error

	FindAllByUser(user *Entity.User) ([]*Entity.Birthday, error)
	FindOneByIdAndUser(id uint, user *Entity.User) (*Entity.Birthday, error)
}
