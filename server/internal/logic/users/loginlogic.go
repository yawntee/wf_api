package users

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"wf_api/server/internal/logic/token"

	"wf_api/server/internal/svc"
	"wf_api/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.Resp, err error) {
	user, err := l.svcCtx.UserModel.FindUser(l.ctx, req.Usr, req.Pwd)
	if err != nil {
		switch err {
		case sqlx.ErrNotFound:
			return &types.Resp{
				Code: 1,
				Msg:  "用户名或密码错误",
			}, nil
		}
		return nil, err
	}
	now := time.Now().Unix()
	expire := l.svcCtx.Config.Auth.AccessExpire
	_token, err := token.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, expire, user.Id)
	if err != nil {
		return nil, err
	}
	return &types.Resp{
		Data: map[string]any{
			"token":  _token,
			"expire": now + expire,
			"refesh": now + expire/2,
		},
	}, nil
}
