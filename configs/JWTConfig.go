package configs

import (
	"os"
	"sync"
)

type JWTConfig struct {
	// Assess token
	ExpiresAt int64 // time in seconds
	Secret    string
}

var jwtConfigOnce sync.Once
var jwtConfig *JWTConfig

func GetJWTConfig() *JWTConfig {
	jwtConfigOnce.Do(func() {
		jwtConfig = &JWTConfig{
			ExpiresAt:/*60 * 10*/ 2, // 10 minutes
			Secret:                  os.Getenv("JWT_TOKEN_SECRET"),
		}
	})

	return jwtConfig
}
