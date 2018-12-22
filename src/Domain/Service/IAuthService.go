package Service

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Domain/Model/ValueObject"
	"github.com/umirode/go-rest/src/Domain/Service/DTO"
)

type IAuthService interface {
	Login(authDTO *DTO.AuthDTO) (*Entity.User, error)
	Signup(authDTO *DTO.AuthDTO) error
	ChangePassword(user *Entity.User, authDTO *DTO.AuthDTO) error

	GetJWTTokenForUser(user *Entity.User, tokenLifeTime int64, tokenSecret string) (*ValueObject.JWTToken, error)
}
