package Service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/umirode/go-rest/src/Domain/Error"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Domain/Model/ValueObject"
	"github.com/umirode/go-rest/src/Domain/Repository"
	"github.com/umirode/go-rest/src/Domain/Service/DTO"
	"time"
)

type AuthService struct {
	userRepository Repository.IUserRepository
}

func NewAuthService(userRepository Repository.IUserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) Login(authDTO DTO.AuthDTO) (*Entity.User, error) {
	user, err := s.userRepository.FindOneByEmailAndPassword(authDTO.Email, s.getPasswordHash(authDTO.Password))
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, Error.NewInvalidError()
	}

	return user, nil
}

func (s *AuthService) Signup(authDTO DTO.AuthDTO) error {
	user, err := s.userRepository.FindOneByEmail(authDTO.Email)
	if err != nil {
		return err
	}

	if user != nil {
		return Error.NewAlreadyExistsError()
	}

	user = &Entity.User{
		Email:        authDTO.Email,
		PasswordHash: s.getPasswordHash(authDTO.Password),
		Birthdays:    *new([]*Entity.Birthday),
	}

	err = s.userRepository.Save(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) ChangePassword(authDTO DTO.AuthDTO) error {
	user, err := s.userRepository.FindOneByEmailAndPassword(authDTO.Email, s.getPasswordHash(authDTO.Password))
	if err != nil {
		return err
	}

	if user == nil {
		return Error.NewInvalidError()
	}

	user.PasswordHash = s.getPasswordHash(authDTO.NewPassword)

	err = s.userRepository.Save(user)
	if err != nil {
		return err
	}

	return nil
}

func (*AuthService) GetJWTTokenForUser(user *Entity.User, tokenLifeTime int64, tokenSecret string) (*ValueObject.JWTToken, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	expiresAt := time.Now().Add(time.Duration(tokenLifeTime) * time.Second).Unix()

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = expiresAt

	t, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return nil, err
	}

	return ValueObject.NewJWTToken(t, expiresAt), nil
}

func (*AuthService) getPasswordHash(password string) string {
	h := sha1.New()

	h.Write([]byte(password))

	return fmt.Sprintf("%x", h.Sum(nil))
}
