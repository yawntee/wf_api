package game

import (
	"context"
	"wf_api/server/internal"

	"wf_api/server/internal/svc"
	"wf_api/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveGameUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveGameUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveGameUserLogic {
	return &RemoveGameUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveGameUserLogic) RemoveGameUser(req *types.SelectGameUserReq) (resp *types.Resp, err error) {
	err = l.svcCtx.GameUserModel.Delete(l.ctx, req.GameUserId)
	if err != nil {
		return nil, err
	}
	return internal.Success("", nil)
}
