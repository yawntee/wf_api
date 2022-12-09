package context

import (
	"github.com/pkg/errors"
	"sync"
)

var (
	ErrAssetUpdate = errors.New("游戏资源更新中。。。")
	UpdateMutex    sync.Mutex
	HttpMutex      sync.Mutex
)
