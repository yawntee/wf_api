package game

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"wf_api/internal"
	"wf_api/internal/svc"
	"wf_api/internal/types"
)

type GameUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGameUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GameUserListLogic {
	return &GameUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GameUserListLogic) GameUserList() (resp *types.Resp, err error) {
	gameUsers, err := l.svcCtx.GameUserModel.GetInfoByUserId(l.ctx, internal.GetUserId(l.ctx))
	if err != nil {
		return nil, err
	}
	return internal.Success("", gameUsers)
}
