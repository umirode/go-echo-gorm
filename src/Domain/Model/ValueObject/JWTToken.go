package ValueObject

import (
	"fmt"
)

type JWTToken struct {
	Token     string
	ExpiresAt int64
}

func (token *JWTToken) ToString() string {
	return fmt.Sprintf("%s | %d", token.Token, token.ExpiresAt)
}

func NewJWTToken(token string, expiresAt int64) *JWTToken {
	return &JWTToken{
		Token:     token,
		ExpiresAt: expiresAt,
	}
}
