package Service

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Domain/Model/ValueObject"
)

type IJWTAuthService interface {
	Refresh(user *Entity.User, token string) (*ValueObject.JWT, *ValueObject.JWT, error)
	CreateByUser(user *Entity.User) (*ValueObject.JWT, *ValueObject.JWT, error)
}
