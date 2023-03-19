package shop

import (
	"context"
	"wf_api/internal"
	"wf_api/internal/svc"
	"wf_api/internal/types"
	"wf_api/wf/api"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShopListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShopListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopListLogic {
	return &ShopListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShopListLogic) ShopList() (resp *types.Resp, err error) {
	if !api.IsInited() {
		return internal.Success("", make([]any, 0))
	}
	return internal.Success("", append(api.EventShops(), api.BossShops()...))
}
