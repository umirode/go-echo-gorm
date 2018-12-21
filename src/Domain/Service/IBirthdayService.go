package Service

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Domain/Service/DTO"
)

type IBirthdayService interface {
	GetAllForUser(user *Entity.User) ([]*Entity.Birthday, error)

	Create(birthdayDTO DTO.BirthdayDTO, user *Entity.User) error
	Update(birthday *Entity.Birthday, birthdayDTO DTO.BirthdayDTO) error
	Delete(birthday *Entity.Birthday) error
}
