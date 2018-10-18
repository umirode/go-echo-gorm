package repositories

import (
	"github.com/stretchr/testify/mock"
	"github.com/umirode/go-rest/models"
)

type MockJWTRefreshTokenRepository struct {
	mock.Mock
}

func (r *MockJWTRefreshTokenRepository) IsUserHasToken(userID uint, token string) (bool, error) {
	args := r.Called(userID, token)

	return args.Bool(0), args.Error(1)
}

func (r *MockJWTRefreshTokenRepository) CountOfTokensAtUser(userID uint) (uint, error) {
	args := r.Called(userID)

	return uint(args.Int(0)), args.Error(1)
}

func (r *MockJWTRefreshTokenRepository) DeleteAllByUser(userID uint) error {
	args := r.Called(userID)

	return args.Error(0)
}

func (r *MockJWTRefreshTokenRepository) DeleteAllByUserAndIP(userID uint, userIP string) error {
	args := r.Called(userID, userIP)

	return args.Error(0)
}

func (r *MockJWTRefreshTokenRepository) AddToken(token *models.JWTRefreshTokenModel) error {
	args := r.Called(token)

	return args.Error(0)
}
