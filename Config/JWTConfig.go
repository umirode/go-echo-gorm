package Config

import (
	"sync"

	"github.com/umirode/go-rest/Config/Helper"
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
			AccessTokenSecret:    Helper.GetEnv("JWT_ACCESS_SECRET", "string").(string),
			AccessTokenLifeTime:  Helper.GetEnv("JWT_ACCESS_LIFE_TIME", "int64").(int64),
			RefreshTokenSecret:   Helper.GetEnv("JWT_REFRESH_SECRET", "string").(string),
			RefreshTokenLifeTime: Helper.GetEnv("JWT_REFRESH_LIFE_TIME", "int64").(int64),
		}
	})

	return jwtConfig
}
