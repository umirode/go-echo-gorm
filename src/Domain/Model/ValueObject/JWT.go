package ValueObject

import (
	"fmt"
)

type JWT struct {
	Token     string
	ExpiresAt int64
}

func (token *JWT) ToString() string {
	return fmt.Sprintf("%s | %d", token.Token, token.ExpiresAt)
}

func NewJWT(token string, expiresAt int64) *JWT {
	return &JWT{
		Token:     token,
		ExpiresAt: expiresAt,
	}
}
