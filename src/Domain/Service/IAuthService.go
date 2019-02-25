package Service

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Domain/Service/DTO"
)

type IAuthService interface {
	Login(authDTO *DTO.AuthDTO) (*Entity.User, error)
	Signup(authDTO *DTO.AuthDTO) (*Entity.User, error)
}
