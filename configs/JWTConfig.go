package configs

import (
	"github.com/umirode/go-rest/Env"
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
			ExpiresAt: Env.GetEnv("JWT_TOKEN_EXPIRES_AT", "int64").(int64),
			Secret:    Env.GetEnv("JWT_TOKEN_SECRET", "string").(string),
		}
	})

	return jwtConfig
}
