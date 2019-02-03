package Config

import (
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
			Host:  GetEnv("SERVER_HOST", "string").(string),
			Port:  GetEnv("SERVER_PORT", "uint").(uint),
			Debug: GetEnv("SERVER_DEBUG", "bool").(bool),
		}
	})

	return serverConfig
}
