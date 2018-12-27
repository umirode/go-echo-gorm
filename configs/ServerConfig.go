package configs

import (
	"github.com/umirode/go-rest/Env"
	"sync"
)

type ServerConfig struct {
	Host  string
	Port  uint
	Debug bool
}

var serverConfigOnce sync.Once
var serverConfig *ServerConfig

func GetServerConfig() *ServerConfig {
	serverConfigOnce.Do(func() {
		serverConfig = &ServerConfig{
			Host:  Env.GetEnv("SERVER_HOST", "string").(string),
			Port:  Env.GetEnv("SERVER_PORT", "uint").(uint),
			Debug: Env.GetEnv("SERVER_DEBUG", "bool").(bool),
		}
	})

	return serverConfig
}
