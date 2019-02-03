package Service

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Domain/Model/ValueObject"
	"github.com/umirode/go-rest/src/Domain/Service/DTO"
)

type IAuthService interface {
	Login(authDTO *DTO.AuthDTO) (*ValueObject.JWT, *ValueObject.JWT, error)
	Signup(authDTO *DTO.AuthDTO) error

	RefreshJWT(user *Entity.User, token string) (*ValueObject.JWT, *ValueObject.JWT, error)
}
