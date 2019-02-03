package Entity

type RefreshToken struct {
	ID uint

	Token     string
	ExpiresAt int64

	OwnerID uint // User
}
