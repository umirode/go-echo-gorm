package Repository

import (
	"time"

	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

type RefreshTokenRepository struct {
	*BaseRepository
}

func NewRefreshTokenRepository() *RefreshTokenRepository {
	return &RefreshTokenRepository{}
}

func (r *RefreshTokenRepository) DeleteOldTokensByUser(user *Entity.User) error {
	r.GetDB().Where("expires_at < ? and owner_id = ?", time.Now().Unix(), user.ID).Delete(&Entity.RefreshToken{})

	return nil
}

func (r *RefreshTokenRepository) Save(token *Entity.RefreshToken) error {
	r.GetDB().Save(token)

	return nil
}

func (r *RefreshTokenRepository) Delete(token *Entity.RefreshToken) error {
	r.GetDB().Delete(token)

	return nil
}

func (r *RefreshTokenRepository) FindOneByTokenAndUser(token string, user *Entity.User) (*Entity.RefreshToken, error) {
	refreshToken := &Entity.RefreshToken{}

	r.GetDB().Where("token = ? and owner_id = ?", token, user.ID).First(refreshToken)
	if r.GetDB().NewRecord(refreshToken) {
		return nil, nil
	}

	return refreshToken, nil
}
