package Module

import (
	"sync"
)

type IModule interface {
	Init(wg *sync.WaitGroup)
	Close(wg *sync.WaitGroup)
}
