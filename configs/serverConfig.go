package configs

import (
	"sync"
)

type ServerConfig struct {
	Host string
	Port uint
}

var serverConfigOnce sync.Once
var serverConfig *ServerConfig

func GetServerConfig() *ServerConfig {
	serverConfigOnce.Do(func() {
		serverConfig = &ServerConfig{
			Host: "127.0.0.1",
			Port: 89,
		}
	})

	return serverConfig
}
