package context

import "github.com/pkg/errors"

var (
	ErrAssetUpdate       = errors.New("游戏资源更新中。。。")
	UpdateMutex    int32 = 0
)
