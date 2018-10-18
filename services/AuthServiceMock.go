package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/mock"
	"github.com/umirode/go-rest/models"
)

type AuthServiceMock struct {
	mock.Mock
}

func (s *AuthServiceMock) Login(email string, password string, ip string, config JWTConfig) (string, string, int64, error) {
	panic("implement me")
}

func (s *AuthServiceMock) DeleteUserRefreshTokensIfMore(user *models.UserModel, count uint) error {
	panic("implement me")
}

func (s *AuthServiceMock) RefreshToken(jwtToken *jwt.Token, config JWTConfig) (string, string, int64, error) {
	panic("implement me")
}

func (s *AuthServiceMock) CreateJWTToken(user *models.UserModel, ip string, secret string, expiresAt int64) (string, int64, error) {
	panic("implement me")
}

func (s *AuthServiceMock) GetPasswordHash(password string) (string, error) {
	panic("implement me")
}

func (s *AuthServiceMock) Signup(email string, password string) error {
	panic("implement me")
}

func (s *AuthServiceMock) Logout(jwtToken *jwt.Token) error {
	panic("implement me")
}

func (s *AuthServiceMock) ResetPassword(jwtToken *jwt.Token, password string, newPassword string) error {
	panic("implement me")
}
