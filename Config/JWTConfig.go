package Config

import (
	"sync"
)

type JWTConfig struct {
	AccessTokenSecret    string
	AccessTokenLifeTime  int64
	RefreshTokenSecret   string
	RefreshTokenLifeTime int64
}

var jwtConfigOnce sync.Once
var jwtConfig *JWTConfig

func GetJWTConfig() *JWTConfig {
	jwtConfigOnce.Do(func() {
		jwtConfig = &JWTConfig{
			AccessTokenSecret:    GetEnv("JWT_ACCESS_SECRET", "string").(string),
			AccessTokenLifeTime:  GetEnv("JWT_ACCESS_LIFE_TIME", "int64").(int64),
			RefreshTokenSecret:   GetEnv("JWT_REFRESH_SECRET", "string").(string),
			RefreshTokenLifeTime: GetEnv("JWT_REFRESH_LIFE_TIME", "int64").(int64),
		}
	})

	return jwtConfig
}
