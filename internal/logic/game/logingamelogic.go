package game

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"strconv"
	"wf_api/internal"
	"wf_api/internal/svc"
	"wf_api/internal/types"
	"wf_api/model"
	"wf_api/wf"
	"wf_api/wf/channel"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginGameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginGameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginGameLogic {
	return &LoginGameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginGameLogic) LoginGame(req *types.LoginGameReq) (resp *types.Resp, err error) {
	_channel, err := channel.ParseChannel(req.Channel)
	if err != nil {
		return nil, err
	}
	c := wf.NewClient(_channel)
	err = c.Login(req.Usr, req.Pwd)
	if err != nil {
		return internal.ReportError(err)
	}
	internal.ClientPool.Set(strconv.FormatUint(c.GameUser.Uid, 10), c, cache.DefaultExpiration)
	data, err := json.Marshal(c)
	if err != nil {
		return internal.ReportError(errors.Wrap(err, "客户端序列化失败"))
	}
	err = l.svcCtx.GameUserModel.UpdateOrInsert(l.ctx, &model.GameUser{
		Id:      c.GameUser.Uid,
		User:    internal.GetUserId(l.ctx),
		Name:    req.Usr,
		Channel: req.Channel,
		Data:    data,
	})
	if err != nil {
		return internal.ReportError(err)
	}
	go func() {
		_ = c.SignUp()
	}()
	return internal.Success(fmt.Sprintf("<%s>登录成功", c.GameUser.Username), map[string]any{
		"id": c.GameUser.Uid,
	})
}
