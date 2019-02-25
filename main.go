package main

import (
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
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

	modules := getModules()
	for _, module := range modules {
		module.Init()
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func getModules() []Module.IModule {
	return []Module.IModule{
		Http.NewModule(),
		Notification.NewModule(),
	}
}
