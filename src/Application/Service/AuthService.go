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
	userRepository         Repository.IUserRepository
	refreshTokenRepository Repository.IRefreshTokenRepository

	accessTokenSecret    string
	accessTokenLifeTime  int64
	refreshTokenSecret   string
	refreshTokenLifeTime int64
}

func NewAuthService(userRepository Repository.IUserRepository, refreshTokenRepository Repository.IRefreshTokenRepository, accessTokenSecret string, accessTokenLifeTime int64, refreshTokenSecret string, refreshTokenLifeTime int64) *AuthService {
	return &AuthService{
		userRepository:         userRepository,
		refreshTokenRepository: refreshTokenRepository,

		accessTokenSecret:    accessTokenSecret,
		accessTokenLifeTime:  accessTokenLifeTime,
		refreshTokenSecret:   refreshTokenSecret,
		refreshTokenLifeTime: refreshTokenLifeTime,
	}
}

func (s *AuthService) Login(authDTO *DTO.AuthDTO) (*ValueObject.JWT, *ValueObject.JWT, error) {
	user, err := s.userRepository.FindOneByEmailAndPassword(authDTO.Email, s.getPasswordHash(authDTO.Password))
	if err != nil {
		return nil, nil, err
	}

	if user == nil {
		return nil, nil, Error.NewInvalidError()
	}

	s.refreshTokenRepository.DeleteOldTokensByUser(user)
	if err != nil {
		return nil, nil, err
	}

	accessToken, err := s.createToken(s.accessTokenSecret, s.accessTokenLifeTime, user.ID)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := s.createToken(s.refreshTokenSecret, s.refreshTokenLifeTime, user.ID)
	if err != nil {
		return nil, nil, err
	}

	err = s.refreshTokenRepository.Save(&Entity.RefreshToken{
		Token:     refreshToken.Token,
		OwnerID:   user.ID,
		ExpiresAt: refreshToken.ExpiresAt,
	})
	if err != nil {
		return nil, nil, err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) RefreshJWT(user *Entity.User, token string) (*ValueObject.JWT, *ValueObject.JWT, error) {
	oldRefreshToken, err := s.refreshTokenRepository.FindOneByTokenAndUser(token, user)
	if err != nil {
		return nil, nil, err
	}

	if oldRefreshToken == nil {
		return nil, nil, Error.NewInvalidError()
	}

	err = s.refreshTokenRepository.Delete(oldRefreshToken)
	if err != nil {
		return nil, nil, err
	}

	s.refreshTokenRepository.DeleteOldTokensByUser(user)
	if err != nil {
		return nil, nil, err
	}

	accessToken, err := s.createToken(s.accessTokenSecret, s.accessTokenLifeTime, user.ID)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := s.createToken(s.refreshTokenSecret, s.refreshTokenLifeTime, user.ID)
	if err != nil {
		return nil, nil, err
	}

	err = s.refreshTokenRepository.Save(&Entity.RefreshToken{
		Token:     refreshToken.Token,
		ExpiresAt: refreshToken.ExpiresAt,
		OwnerID:   user.ID,
	})
	if err != nil {
		return nil, nil, err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) Signup(authDTO *DTO.AuthDTO) error {
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
	}

	err = s.userRepository.Save(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) createToken(secret string, lifeTime int64, userID uint) (*ValueObject.JWT, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	expiresAt := time.Now().Add(time.Duration(lifeTime) * time.Second).Unix()

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = expiresAt

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return ValueObject.NewJWT(t, expiresAt), nil
}

func (*AuthService) getPasswordHash(password string) string {
	h := sha1.New()

	h.Write([]byte(password))

	return fmt.Sprintf("%x", h.Sum(nil))
}
