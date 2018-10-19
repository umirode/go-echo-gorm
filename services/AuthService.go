package services

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/umirode/go-rest/errors"
	"github.com/umirode/go-rest/models"
	"github.com/umirode/go-rest/repositories"
	"time"
)

type AuthService struct {
	BaseService

	UserRepository            repositories.IUserRepository
	JWTRefreshTokenRepository repositories.IJWTRefreshTokenRepository
}

func (s *AuthService) DeleteUserByEmail(email string) error {
	err := s.UserRepository.DeleteUserByEmail(email)

	return err
}

func (s *AuthService) Login(email string, password string, userIP string, config JWTConfig) (string, string, int64, error) {
	// Get password hash
	passwordHash, err := s.GetPasswordHash(password)
	if err != nil {
		return "", "", 0, err
	}

	// Find user in database
	user, err := s.UserRepository.FindSingleByEmailAndPassword(email, passwordHash)
	if err != nil {
		return "", "", 0, errors.NewInvalidEmailOrPasswordError()
	}

	// Create assess token
	token, expiresAt, err := s.CreateJWTToken(user.ID, userIP, config.Secret, config.ExpiresAt)
	if err != nil {
		return "", "", 0, err
	}

	// Create refresh token
	refreshToken, _, err := s.CreateJWTToken(user.ID, userIP, config.RefreshSecret, config.RefreshExpiresAt)
	if err != nil {
		return "", "", 0, err
	}

	err = s.JWTRefreshTokenRepository.DeleteAllByUserAndIP(user.ID, userIP)
	if err != nil {
		return "", "", 0, err
	}

	// Check refresh tokens in database and delete them if more than 10
	err = s.DeleteUserRefreshTokensIfMore(user.ID, 10)
	if err != nil {
		return "", "", 0, err
	}

	// Add new refresh token to database
	err = s.JWTRefreshTokenRepository.AddToken(&models.JWTRefreshTokenModel{
		UserIP: userIP,
		UserID: user.ID,
		Token:  refreshToken,
	})
	if err != nil {
		return "", "", 0, err
	}

	return token, refreshToken, expiresAt, nil
}

func (s *AuthService) DeleteUserRefreshTokensIfMoreByEmail(email string, count uint) error {
	user, err := s.UserRepository.FindSingleByEmail(email)
	if err != nil {
		return errors.NewAuthError()
	}

	return s.DeleteUserRefreshTokensIfMore(user.ID, count)
}

func (s *AuthService) DeleteUserRefreshTokensIfMore(userID uint, count uint) error {
	// Get user refresh tokens count
	userRefreshTokenCount, err := s.JWTRefreshTokenRepository.CountOfTokensAtUser(userID)
	if err != nil {
		return err
	}

	// Reset all user refresh tokens if more then COUNT
	if userRefreshTokenCount >= count {
		err := s.JWTRefreshTokenRepository.DeleteAllByUser(userID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *AuthService) RefreshToken(userID uint, userIP string, token string, config JWTConfig) (string, string, int64, error) {
	// Find user in database
	user, err := s.UserRepository.FindSingleByID(userID)
	if err != nil {
		return "", "", 0, errors.NewAuthError()
	}

	// Check if user has current refresh token in database
	userHasToken, err := s.JWTRefreshTokenRepository.IsUserHasToken(user.ID, token)
	if err != nil || !userHasToken {
		return "", "", 0, errors.NewAuthError()
	}

	err = s.JWTRefreshTokenRepository.DeleteAllByUserAndIP(user.ID, userIP)
	if err != nil {
		return "", "", 0, err
	}

	// Reset all user refresh tokens if more then COUNT
	err = s.DeleteUserRefreshTokensIfMore(user.ID, 10)
	if err != nil {
		return "", "", 0, err
	}

	// Create assess token
	token, expiresAt, err := s.CreateJWTToken(user.ID, userIP, config.Secret, config.ExpiresAt)
	if err != nil {
		return "", "", 0, errors.NewAuthError()
	}

	// Create refresh token
	refreshToken, _, err := s.CreateJWTToken(user.ID, userIP, config.RefreshSecret, config.RefreshExpiresAt)
	if err != nil {
		return "", "", 0, errors.NewAuthError()
	}

	// Add new refresh token to database
	err = s.JWTRefreshTokenRepository.AddToken(&models.JWTRefreshTokenModel{
		UserIP: userIP,
		UserID: user.ID,
		Token:  refreshToken,
	})
	if err != nil {
		return "", "", 0, err
	}

	return token, refreshToken, expiresAt, err
}

func (s *AuthService) CreateJWTToken(userID uint, userIP string, secret string, expiresAt int64) (string, int64, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	expiresAt = time.Now().Add(time.Duration(expiresAt) * time.Second).Unix()

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["user_ip"] = userIP
	claims["exp"] = expiresAt

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}

	return t, expiresAt, nil
}

func (s *AuthService) GetPasswordHash(password string) (string, error) {
	h := sha1.New()

	h.Write([]byte(password))

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func (s *AuthService) Signup(email string, password string) error {
	// Find user in database
	user, err := s.UserRepository.FindSingleByEmail(email)
	if user != nil {
		return errors.NewAlreadyExistsError()
	}

	passwordHash, err := s.GetPasswordHash(password)
	if err != nil {
		return err
	}

	err = s.UserRepository.AddUser(&models.UserModel{
		Email:    email,
		Password: passwordHash,
	})

	return err
}

func (s *AuthService) Logout(userID uint, userIP string) error {
	user, err := s.UserRepository.FindSingleByID(userID)
	if err != nil {
		return errors.NewAuthError()
	}

	return s.JWTRefreshTokenRepository.DeleteAllByUserAndIP(user.ID, userIP)
}

func (s *AuthService) ResetPassword(userID uint, password string, newPassword string) error {
	// Find user in database
	user, err := s.UserRepository.FindSingleByID(userID)
	if err != nil {
		return errors.NewAuthError()
	}

	passwordHash, err := s.GetPasswordHash(password)
	if err != nil {
		return err
	}

	newPasswordHash, err := s.GetPasswordHash(newPassword)
	if err != nil {
		return err
	}

	user, err = s.UserRepository.FindSingleByEmailAndPassword(user.Email, passwordHash)
	if err != nil || user == nil {
		return errors.NewInvalidOldPasswordError()
	}

	user.Password = newPasswordHash
	err = s.UserRepository.UpdateUser(user)
	if err != nil {
		return err
	}

	err = s.DeleteUserRefreshTokensIfMore(user.ID, 0)
	if err != nil {
		return err
	}

	return nil
}
