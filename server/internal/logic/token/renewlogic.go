package token

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"wf_api/server/internal"
	"wf_api/server/internal/svc"
	"wf_api/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RenewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRenewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RenewLogic {
	return &RenewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func GetJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *RenewLogic) Renew() (resp *types.Resp, err error) {
	now := time.Now().Unix()
	expire := l.svcCtx.Config.Auth.AccessExpire
	_token, err := GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, expire, internal.GetUserId(l.ctx))
	if err != nil {
		return nil, err
	}
	return &types.Resp{
		Msg: "登录成功",
		Data: map[string]any{
			"token":  _token,
			"expire": now + expire,
			"refesh": now + expire/2,
		},
	}, nil
}
