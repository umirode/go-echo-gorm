package repositories

import (
	"github.com/umirode/go-rest/models"
)

type IJWTRefreshTokenRepository interface {
	IsUserHasToken(userID uint, token string) (bool, error)
	CountOfTokensAtUser(userID uint) (uint, error)
	DeleteAllByUser(userID uint) error
	DeleteAllByUserAndIP(userID uint, userIP string) error
	AddToken(token *models.JWTRefreshTokenModel) error
}
