package configs

import (
	"os"
	"strconv"
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
		port, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
		debug, _ := strconv.ParseBool(os.Getenv("SERVER_DEBUG"))

		serverConfig = &ServerConfig{
			Host:  os.Getenv("SERVER_HOST"),
			Port:  uint(port),
			Debug: debug,
		}
	})

	return serverConfig
}
