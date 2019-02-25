package Service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/umirode/go-rest/src/Domain/Error"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Domain/Model/ValueObject"
	"github.com/umirode/go-rest/src/Domain/Repository"
)

type JWTAuthService struct {
	refreshTokenRepository Repository.IRefreshTokenRepository

	accessTokenSecret    string
	accessTokenLifeTime  int64
	refreshTokenSecret   string
	refreshTokenLifeTime int64
}

func NewJWTAuthService(
	refreshTokenRepository Repository.IRefreshTokenRepository,
	accessTokenSecret string,
	accessTokenLifeTime int64,
	refreshTokenSecret string,
	refreshTokenLifeTime int64,
) *JWTAuthService {
	return &JWTAuthService{
		refreshTokenRepository: refreshTokenRepository,

		accessTokenSecret:    accessTokenSecret,
		accessTokenLifeTime:  accessTokenLifeTime,
		refreshTokenSecret:   refreshTokenSecret,
		refreshTokenLifeTime: refreshTokenLifeTime,
	}
}

func (s *JWTAuthService) Refresh(user *Entity.User, token string) (*ValueObject.JWT, *ValueObject.JWT, error) {
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

	err = s.refreshTokenRepository.DeleteOldTokensByUser(user)
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

func (s *JWTAuthService) CreateByUser(user *Entity.User) (*ValueObject.JWT, *ValueObject.JWT, error) {
	err := s.refreshTokenRepository.DeleteOldTokensByUser(user)
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
		OwnerID:   user.ID,
		Token:     refreshToken.Token,
		ExpiresAt: refreshToken.ExpiresAt,
	})
	if err != nil {
		return nil, nil, err
	}

	return accessToken, refreshToken, nil
}

func (s *JWTAuthService) createToken(secret string, lifeTime int64, userID uint) (*ValueObject.JWT, error) {
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
