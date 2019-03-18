package Http

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/umirode/go-rest/Config"
	"github.com/umirode/go-rest/Module/Http/Router"
)

type Module struct {
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Init(wg *sync.WaitGroup) {
	go func() {
		/**
		Get server address
		*/
		serverConfig := Config.GetServerConfig()
		serverAddress := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)

		/**
		Start server
		*/
		logrus.Fatal(Router.NewRouter(serverConfig.Debug).Router.Start(serverAddress))
	}()
}

func (m *Module) Close(wg *sync.WaitGroup) {
	println("test")
}
