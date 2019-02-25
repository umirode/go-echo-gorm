package Notification

import (
	"github.com/umirode/go-rest/Config"
)

type Module struct {
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Init() {
	go func() {
		firebaseConfig := Config.GetFirebaseConfig()

		notification := NewHandler(firebaseConfig.CloudMessagingKey)
		notification.Run()
	}()
}
