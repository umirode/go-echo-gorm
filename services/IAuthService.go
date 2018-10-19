package services

type JWTConfig struct {
	// Assess token
	ExpiresAt int64 // time in seconds
	Secret    string

	// Refresh token
	RefreshExpiresAt int64 // time in seconds
	RefreshSecret    string
}

type IAuthService interface {
	Login(email string, password string, userIP string, config JWTConfig) (string, string, int64, error)
	DeleteUserRefreshTokensIfMore(userID uint, count uint) error
	DeleteUserRefreshTokensIfMoreByEmail(email string, count uint) error
	DeleteUserByEmail(email string) error
	RefreshToken(userID uint, userIP string, token string, config JWTConfig) (string, string, int64, error)
	CreateJWTToken(userID uint, userIP string, secret string, expiresAt int64) (string, int64, error)
	GetPasswordHash(password string) (string, error)
	Signup(email string, password string) error
	Logout(userID uint, userIP string) error
	ResetPassword(userID uint, password string, newPassword string) error
}
