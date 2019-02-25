package Service

import (
	"crypto/sha1"
	"fmt"

	"github.com/umirode/go-rest/src/Domain/Error"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Domain/Repository"
	"github.com/umirode/go-rest/src/Domain/Service/DTO"
)

type AuthService struct {
	userRepository Repository.IUserRepository
}

func NewAuthService(
	userRepository Repository.IUserRepository,
) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) Login(authDTO *DTO.AuthDTO) (*Entity.User, error) {
	user, err := s.userRepository.FindOneByEmailAndPassword(authDTO.Email, s.getPasswordHash(authDTO.Password))
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, Error.NewInvalidError()
	}

	return user, nil
}

func (s *AuthService) Signup(authDTO *DTO.AuthDTO) (*Entity.User, error) {
	user, err := s.userRepository.FindOneByEmail(authDTO.Email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, Error.NewAlreadyExistsError()
	}

	user = &Entity.User{
		Email:        authDTO.Email,
		PasswordHash: s.getPasswordHash(authDTO.Password),
	}

	err = s.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (*AuthService) getPasswordHash(password string) string {
	h := sha1.New()

	h.Write([]byte(password))

	return fmt.Sprintf("%x", h.Sum(nil))
}
