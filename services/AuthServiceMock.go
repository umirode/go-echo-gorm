package services

import (
	"github.com/stretchr/testify/mock"
)

type AuthServiceMock struct {
	mock.Mock
}

func (s *AuthServiceMock) DeleteUserRefreshTokensIfMore(userID uint, count uint) error {
	panic("implement me")
}

func (s *AuthServiceMock) DeleteUserRefreshTokensIfMoreByEmail(email string, count uint) error {
	panic("implement me")
}

func (s *AuthServiceMock) DeleteUserByEmail(email string) error {
	panic("implement me")
}

func (s *AuthServiceMock) CreateJWTToken(userID uint, userIP string, secret string, expiresAt int64) (string, int64, error) {
	panic("implement me")
}

func (s *AuthServiceMock) RefreshToken(userID uint, userIP string, token string, config JWTConfig) (string, string, int64, error) {
	panic("implement me")
}

func (s *AuthServiceMock) Logout(userID uint, userIP string) error {
	panic("implement me")
}

func (s *AuthServiceMock) ChangePassword(userID uint, password string, newPassword string) error {
	panic("implement me")
}

func (s *AuthServiceMock) Login(email string, password string, ip string, config JWTConfig) (string, string, int64, error) {
	panic("implement me")
}

func (s *AuthServiceMock) GetPasswordHash(password string) (string, error) {
	panic("implement me")
}

func (s *AuthServiceMock) Signup(email string, password string) error {
	panic("implement me")
}
