package Repository

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

type IRefreshTokenRepository interface {
	Save(token *Entity.RefreshToken) error
	Delete(token *Entity.RefreshToken) error

	CountByUser(user *Entity.User) (uint, error)
	FindOneByTokenAndUser(token string, user *Entity.User) (*Entity.RefreshToken, error)
	DeleteOldTokensByUser(user *Entity.User) error
}
