package Entity

type RefreshToken struct {
	ID uint `gorm:"primary_key"`

	Token     string `gorm:"size:255"`
	ExpiresAt int64  `gorm:"size:20"`

	OwnerID uint // User
}
