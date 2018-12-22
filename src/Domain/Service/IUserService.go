package Service

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

type IUserService interface {
	GetOneById(id uint) (*Entity.User, error)
}
