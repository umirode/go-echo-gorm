package Service

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Domain/Service/DTO"
)

type IBirthdayService interface {
	GetAllByUser(user *Entity.User) ([]*Entity.Birthday, error)
	GetOneByIdAndUser(id uint, user *Entity.User) (*Entity.Birthday, error)

	Create(birthdayDTO *DTO.BirthdayDTO, user *Entity.User) error
	Update(birthday *Entity.Birthday, birthdayDTO *DTO.BirthdayDTO, user *Entity.User) error
	Delete(birthday *Entity.Birthday, user *Entity.User) error
}
