package internal

import (
	"context"
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"strconv"
	"time"
	model "wf_api/model"
	"wf_api/util"
	"wf_api/wf"
)

type _ClientPool cache.Cache

var ClientPool = (*_ClientPool)(cache.New(5*time.Minute, 10*time.Minute))

var ErrInvalidGameUserId = errors.New("无效的游戏用户ID")

func (c *_ClientPool) GetClient(ctx context.Context, gameUserModel model.GameUserModel, gameUserId int64) (*wf.Client, error) {
	id := strconv.FormatInt(gameUserId, 10)
	if v, ok := c.Get(id); ok {
		return v.(*wf.Client), nil
	}
	data, err := gameUserModel.GetData(ctx, gameUserId)
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
	err = client.CheckLogin()
	if err != nil {
		return nil, err
	}
	err = gameUserModel.Update(ctx, &model.GameUser{
		Id:      client.GameUser.Uid,
		User:    GetUserId(ctx),
		Channel: uint8(client.Channel.Id),
		Name:    client.GameUser.Username,
		Data:    util.ToJson(client),
	})
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

var OtpCache = cache.New(5*time.Minute, 10*time.Minute)
