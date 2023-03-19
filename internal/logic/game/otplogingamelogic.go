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

	"github.com/zeromicro/go-zero/core/logx"
)

type OtpLoginGameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOtpLoginGameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OtpLoginGameLogic {
	return &OtpLoginGameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OtpLoginGameLogic) OtpLoginGame(req *types.OtpLoginGameReq) (resp *types.Resp, err error) {
	userId := internal.GetUserId(l.ctx)
	_c, ok := internal.OtpCache.Get(strconv.FormatInt(userId, 10))
	if !ok {
		return internal.ReportMsg("您还未发送验证码")
	}
	c := _c.(*wf.Client)
	err = c.OtpLogin(req.Phone, req.Otp)
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
		User:    userId,
		Name:    req.Phone,
		Channel: uint8(c.Channel.Id),
		Data:    data,
	})
	if err != nil {
		return internal.ReportError(err)
	}
	return internal.Success(fmt.Sprintf("<%s>登录成功", c.GameUser.Username), map[string]any{
		"id": c.GameUser.Uid,
	})
}
