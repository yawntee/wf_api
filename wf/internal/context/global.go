package context

import (
	"sync"
)

var (
	UpdateMutex sync.Mutex
	HttpMutex   sync.Mutex
)
