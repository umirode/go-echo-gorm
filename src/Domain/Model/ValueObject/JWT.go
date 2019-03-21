package ValueObject

type JWT struct {
	Token     string
	ExpiresAt int64
}

func NewJWT(token string, expiresAt int64) *JWT {
	return &JWT{
		Token:     token,
		ExpiresAt: expiresAt,
	}
}
