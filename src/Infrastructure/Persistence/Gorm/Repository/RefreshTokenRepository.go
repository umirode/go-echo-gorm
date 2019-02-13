package Repository

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"time"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		db: db,
	}
}

func (r *RefreshTokenRepository) DeleteOldTokensByUser(user *Entity.User) error {
	r.db.Where("expires_at < ? and owner_id = ?", time.Now().Unix(), user.ID).Delete(&Entity.RefreshToken{})

	return nil
}

func (r *RefreshTokenRepository) Save(token *Entity.RefreshToken) error {
	r.db.Save(token)

	return nil
}

func (r *RefreshTokenRepository) Delete(token *Entity.RefreshToken) error {
	r.db.Delete(token)

	return nil
}

func (r *RefreshTokenRepository) FindOneByTokenAndUser(token string, user *Entity.User) (*Entity.RefreshToken, error) {
	refreshToken := &Entity.RefreshToken{}

	r.db.Where("token = ? and owner_id = ?", token, user.ID).First(refreshToken)
	if r.db.NewRecord(refreshToken) {
		return nil, nil
	}

	return refreshToken, nil
}
