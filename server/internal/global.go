package internal

import (
	"context"
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"strconv"
	"time"
	"wf_api/server/model"
	"wf_api/server/wf"
)

type _ClientPool cache.Cache

var ClientPool = (*_ClientPool)(cache.New(5*time.Minute, 10*time.Minute))

var ErrInvalidGameUserId = errors.New("无效的游戏用户ID")

func (c *_ClientPool) GetClient(ctx context.Context, model model.GameUserModel, gameUserId int64) (*wf.Client, error) {
	id := strconv.FormatInt(gameUserId, 10)
	if v, ok := c.Get(id); ok {
		return v.(*wf.Client), nil
	}
	data, err := model.GetData(ctx, gameUserId)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, ErrInvalidGameUserId
		}
		return nil, err
	}
	var client wf.Client
	err = json.Unmarshal(data, &client)
	if err != nil {
		return nil, err
	}
	err = client.SignUp()
	if err != nil {
		return nil, err
	}
	c.Set(id, &client, cache.DefaultExpiration)
	return &client, nil
}

type TaskResult struct {
	Id  int64
	Msg string
}

var TaskResults = make(map[int64]chan TaskResult)

var TaskMutex = make(map[int64]*int32)

var OtpCache cache.Cache
