package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/umirode/go-rest/models"
)

type JWTConfig struct {
	// Assess token
	ExpiresAt int64 // time in seconds
	Secret    string

	// Refresh token
	RefreshExpiresAt int64 // time in seconds
	RefreshSecret    string
}

type IAuthService interface {
	Login(email string, password string, ip string, config JWTConfig) (string, string, int64, error)
	DeleteUserRefreshTokensIfMore(user *models.UserModel, count uint) error
	RefreshToken(jwtToken *jwt.Token, config JWTConfig) (string, string, int64, error)
	CreateJWTToken(user *models.UserModel, ip string, secret string, expiresAt int64) (string, int64, error)
	GetPasswordHash(password string) (string, error)
	Signup(email string, password string) error
	Logout(jwtToken *jwt.Token) error
	ResetPassword(jwtToken *jwt.Token, password string, newPassword string) error
}
