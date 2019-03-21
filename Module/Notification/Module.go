package Notification

import (
	"sync"

	"github.com/umirode/go-rest/Config"
)

type Module struct {
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Init(wg *sync.WaitGroup) {
	go func() {
		firebaseConfig := Config.GetFirebaseConfig()

		notification := NewHandler(firebaseConfig.CloudMessagingKey)
		notification.Run()
	}()
}

func (m *Module) Close(wg *sync.WaitGroup) {

}
