package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/umirode/go-rest/models"
	"github.com/umirode/go-rest/repositories"
	"testing"
)

func TestAuthService_GetPasswordHash(t *testing.T) {
	userRepository := &repositories.MockUserRepository{}
	jwtRefreshTokenRepository := &repositories.MockJWTRefreshTokenRepository{}

	s := &AuthService{
		UserRepository:            userRepository,
		JWTRefreshTokenRepository: jwtRefreshTokenRepository,
	}

	hash, err := s.GetPasswordHash("test")

	assert.NotEmpty(t, hash)
	assert.NoError(t, err)
}

func TestAuthService_CreateJWTToken(t *testing.T) {
	userRepository := &repositories.MockUserRepository{}
	jwtRefreshTokenRepository := &repositories.MockJWTRefreshTokenRepository{}

	s := &AuthService{
		UserRepository:            userRepository,
		JWTRefreshTokenRepository: jwtRefreshTokenRepository,
	}

	token, expiresAt, err := s.CreateJWTToken(&models.UserModel{}, "test", "test", 1000)

	assert.NotEmpty(t, token)
	assert.NotEmpty(t, expiresAt)
	assert.NoError(t, err)
}

func TestAuthService_DeleteUserRefreshTokensIfMore(t *testing.T) {
	userRepository := &repositories.MockUserRepository{}
	jwtRefreshTokenRepository := &repositories.MockJWTRefreshTokenRepository{}

	jwtRefreshTokenRepository.On("CountOfTokensAtUser", mock.AnythingOfType("uint")).Return(10, nil)
	jwtRefreshTokenRepository.On("DeleteAllByUser", mock.AnythingOfType("uint")).Return(nil)

	s := &AuthService{
		UserRepository:            userRepository,
		JWTRefreshTokenRepository: jwtRefreshTokenRepository,
	}

	err := s.DeleteUserRefreshTokensIfMore(&models.UserModel{}, 1)

	assert.NoError(t, err)
}

func TestAuthService_Login(t *testing.T) {
	userRepository := &repositories.MockUserRepository{}

	userRepository.On("FindSingleByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&models.UserModel{}, nil)

	jwtRefreshTokenRepository := &repositories.MockJWTRefreshTokenRepository{}

	jwtRefreshTokenRepository.On("CountOfTokensAtUser", mock.AnythingOfType("uint")).Return(10, nil)
	jwtRefreshTokenRepository.On("DeleteAllByUser", mock.AnythingOfType("uint")).Return(nil)
	jwtRefreshTokenRepository.On("DeleteAllByUserAndIP", mock.AnythingOfType("uint"), mock.AnythingOfType("string")).Return(nil)
	jwtRefreshTokenRepository.On("AddToken", mock.Anything).Return(nil)

	s := &AuthService{
		UserRepository:            userRepository,
		JWTRefreshTokenRepository: jwtRefreshTokenRepository,
	}

	assessToken, refreshToken, expiresAt, err := s.Login("test", "test", "test", JWTConfig{})

	assert.NotEmpty(t, assessToken)
	assert.NotEmpty(t, refreshToken)
	assert.NotEmpty(t, expiresAt)
	assert.NoError(t, err)
}

func TestAuthService_Logout(t *testing.T) {
	userRepository := &repositories.MockUserRepository{}

	userRepository.On("FindSingleByID", mock.AnythingOfType("uint")).Return(&models.UserModel{}, nil)

	jwtRefreshTokenRepository := &repositories.MockJWTRefreshTokenRepository{}

	jwtRefreshTokenRepository.On("DeleteAllByUserAndIP", mock.AnythingOfType("uint"), mock.AnythingOfType("string")).Return(nil)

	s := &AuthService{
		UserRepository:            userRepository,
		JWTRefreshTokenRepository: jwtRefreshTokenRepository,
	}

	err := s.Logout(&jwt.Token{
		Claims: jwt.MapClaims{
			"user_id": float64(1),
			"user_ip": "test",
		},
	})

	assert.NoError(t, err)
}

func TestAuthService_RefreshToken(t *testing.T) {
	userRepository := &repositories.MockUserRepository{}

	userRepository.On("FindSingleByID", mock.AnythingOfType("uint")).Return(&models.UserModel{}, nil)

	jwtRefreshTokenRepository := &repositories.MockJWTRefreshTokenRepository{}

	jwtRefreshTokenRepository.On("IsUserHasToken", mock.AnythingOfType("uint"), mock.AnythingOfType("string")).Return(true, nil)
	jwtRefreshTokenRepository.On("CountOfTokensAtUser", mock.AnythingOfType("uint")).Return(10, nil)
	jwtRefreshTokenRepository.On("DeleteAllByUser", mock.AnythingOfType("uint")).Return(nil)
	jwtRefreshTokenRepository.On("DeleteAllByUserAndIP", mock.AnythingOfType("uint"), mock.AnythingOfType("string")).Return(nil)
	jwtRefreshTokenRepository.On("AddToken", mock.Anything).Return(nil)

	s := &AuthService{
		UserRepository:            userRepository,
		JWTRefreshTokenRepository: jwtRefreshTokenRepository,
	}

	assessToken, refreshToken, expiresAt, err := s.RefreshToken(&jwt.Token{
		Claims: jwt.MapClaims{
			"user_id": float64(1),
			"user_ip": "test",
		},
	}, JWTConfig{})

	assert.NotEmpty(t, assessToken)
	assert.NotEmpty(t, refreshToken)
	assert.NotEmpty(t, expiresAt)
	assert.NoError(t, err)
}

func TestAuthService_ResetPassword(t *testing.T) {
	userRepository := &repositories.MockUserRepository{}

	userRepository.On("FindSingleByID", mock.AnythingOfType("uint")).Return(&models.UserModel{}, nil)
	userRepository.On("FindSingleByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&models.UserModel{}, nil)
	userRepository.On("UpdateUser", mock.Anything).Return(nil)

	jwtRefreshTokenRepository := &repositories.MockJWTRefreshTokenRepository{}

	jwtRefreshTokenRepository.On("CountOfTokensAtUser", mock.AnythingOfType("uint")).Return(10, nil)
	jwtRefreshTokenRepository.On("DeleteAllByUser", mock.AnythingOfType("uint")).Return(nil)

	s := &AuthService{
		UserRepository:            userRepository,
		JWTRefreshTokenRepository: jwtRefreshTokenRepository,
	}

	err := s.ResetPassword(&jwt.Token{
		Claims: jwt.MapClaims{
			"user_id": float64(1),
			"user_ip": "test",
		},
	}, "test", "test")

	assert.NoError(t, err)
}

func TestAuthService_Signup(t *testing.T) {
	userRepository := &repositories.MockUserRepository{}

	userRepository.On("AddUser", mock.Anything).Return(nil)
	userRepository.On("FindSingleByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, nil)

	jwtRefreshTokenRepository := &repositories.MockJWTRefreshTokenRepository{}

	s := &AuthService{
		UserRepository:            userRepository,
		JWTRefreshTokenRepository: jwtRefreshTokenRepository,
	}

	err := s.Signup("test", "test")

	assert.NoError(t, err)
}
