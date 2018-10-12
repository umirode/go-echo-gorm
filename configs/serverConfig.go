package configs

import (
	"os"
	"strconv"
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
		port, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))

		serverConfig = &ServerConfig{
			Host: os.Getenv("SERVER_HOST"),
			Port: uint(port),
		}
	})

	return serverConfig
}
