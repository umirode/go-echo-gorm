package ValueObject

import (
	"fmt"
)

type JWTToken struct {
	token     string
	expiresAt int64
}

func (token *JWTToken) ToString() string {
	return fmt.Sprintf("%s | %d", token.token, token.expiresAt)
}

func NewJWTToken(token string, expiresAt int64) *JWTToken {
	return &JWTToken{
		token:     token,
		expiresAt: expiresAt,
	}
}
