package Config

import (
	"sync"

	"github.com/umirode/go-rest/Config/Helper"
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
			Host:  Helper.GetEnv("SERVER_HOST", "string").(string),
			Port:  Helper.GetEnv("SERVER_PORT", "uint").(uint),
			Debug: Helper.GetEnv("SERVER_DEBUG", "bool").(bool),
		}
	})

	return serverConfig
}
