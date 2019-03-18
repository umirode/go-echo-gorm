package main

import (
	"sync"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/umirode/go-rest/Module"
	"github.com/umirode/go-rest/Module/Http"
	"github.com/umirode/go-rest/Module/Notification"
)

func main() {
	/**
	Load .env variables
	*/
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err)
	}

	wg := &sync.WaitGroup{}

	modules := getModules()
	wg.Add(len(modules))

	for _, module := range modules {
		module.Init(wg)
	}

	wg.Wait()

	for _, module := range modules {
		module.Close(wg)
	}
}

func getModules() []Module.IModule {
	return []Module.IModule{
		Http.NewModule(),
		Notification.NewModule(),
	}
}
