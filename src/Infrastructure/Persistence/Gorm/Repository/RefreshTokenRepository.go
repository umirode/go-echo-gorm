package Repository

import (
	"time"

	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

type RefreshTokenRepository struct {
	BaseRepository
}

func NewRefreshTokenRepository() *RefreshTokenRepository {
	return &RefreshTokenRepository{}
}

func (r *RefreshTokenRepository) DeleteOldTokensByUser(user *Entity.User) error {
	r.GetGormDB().Where("expires_at < ? and owner_id = ?", time.Now().Unix(), user.ID).Delete(&Entity.RefreshToken{})

	return nil
}

func (r *RefreshTokenRepository) DeleteAllTokensByUser(user *Entity.User) error {
	r.GetGormDB().Where("owner_id = ?", user.ID).Delete(&Entity.RefreshToken{})

	return nil
}

func (r *RefreshTokenRepository) Save(token *Entity.RefreshToken) error {
	r.GetGormDB().Save(token)

	return nil
}

func (r *RefreshTokenRepository) Delete(token *Entity.RefreshToken) error {
	r.GetGormDB().Delete(token)

	return nil
}

func (r *RefreshTokenRepository) CountByUser(user *Entity.User) (uint, error) {
	count := new(uint)

	r.GetGormDB().Model(&Entity.RefreshToken{}).Where("owner_id = ?", user.ID).Count(count)

	return *count, nil
}

func (r *RefreshTokenRepository) FindOneByTokenAndUser(token string, user *Entity.User) (*Entity.RefreshToken, error) {
	refreshToken := &Entity.RefreshToken{}

	r.GetGormDB().Where("token = ? and owner_id = ?", token, user.ID).First(refreshToken)
	if r.GetGormDB().NewRecord(refreshToken) {
		return nil, nil
	}

	return refreshToken, nil
}
