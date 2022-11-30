package game

import (
	"context"
	"fmt"
	"github.com/patrickmn/go-cache"
	"strconv"
	"sync"
	"wf_api/server/internal"
	"wf_api/server/wf"
	"wf_api/server/wf/channel"

	"wf_api/server/internal/svc"
	"wf_api/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendLoginGameOtpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendLoginGameOtpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLoginGameOtpLogic {
	return &SendLoginGameOtpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var mu sync.Mutex

func (l *SendLoginGameOtpLogic) SendLoginGameOtp(req *types.SendOtpReq) (resp *types.Resp, err error) {
	userId := strconv.FormatInt(internal.GetUserId(l.ctx), 10)
	if _, expire, ok := internal.OtpCache.GetWithExpiration(userId); ok {
		return internal.ReportMsg(fmt.Sprintf("发送冷却中，请%d秒后再进行进行操作", expire.Second()))
	}
	defer func() {
		if v, ok := internal.OtpCache.Get(userId); ok && v == nil {
			internal.OtpCache.Delete(userId)
		}
	}()
	internal.OtpCache.Set(userId, nil, cache.DefaultExpiration)
	_channel, err := channel.ParseChannel(req.Channel)
	if err != nil {
		return nil, err
	}
	c := wf.NewClient(_channel)
	err = c.SendOtp(req.Phone)
	if err != nil {
		return internal.ReportError(err)
	}
	internal.OtpCache.Set(userId, c, cache.DefaultExpiration)
	return internal.Success("", nil)
}
