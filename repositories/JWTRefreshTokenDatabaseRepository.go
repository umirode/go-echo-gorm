package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/go-rest/errors"
	"github.com/umirode/go-rest/models"
)

type JWTRefreshTokenDatabaseRepository struct {
	BaseDatabaseRepository
}

func NewJWTRefreshTokenDatabaseRepository(database *gorm.DB) *JWTRefreshTokenDatabaseRepository {
	repository := &JWTRefreshTokenDatabaseRepository{}

	repository.Database = database

	return repository
}

func (r *JWTRefreshTokenDatabaseRepository) IsUserHasToken(userID uint, token string) (bool, error) {
	jwtRefreshToken := new(models.JWTRefreshTokenModel)

	result := r.Database.Where("user_id = ? AND token = ?", userID, token).First(&jwtRefreshToken)
	if result.Error != nil {
		return false, result.Error
	}

	if jwtRefreshToken.ID == 0 {
		return false, errors.NewNotFoundError()
	}

	return true, nil
}

func (r *JWTRefreshTokenDatabaseRepository) CountOfTokensAtUser(userID uint) (uint, error) {
	count := 0

	result := r.Database.Model(&models.JWTRefreshTokenModel{}).Where("user_id = ?", userID).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return uint(count), nil
}

func (r *JWTRefreshTokenDatabaseRepository) DeleteAllByUser(userID uint) error {
	result := r.Database.Where("user_id = ?", userID).Delete(&models.JWTRefreshTokenModel{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *JWTRefreshTokenDatabaseRepository) DeleteAllByUserAndIP(userID uint, userIP string) error {
	result := r.Database.Where("user_id = ? AND user_ip = ?", userID, userIP).Delete(&models.JWTRefreshTokenModel{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *JWTRefreshTokenDatabaseRepository) AddToken(token *models.JWTRefreshTokenModel) error {
	err := r.create(token)

	return err
}
