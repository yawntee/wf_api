package shop

import (
	"context"
	"wf_api/server/internal"
	"wf_api/server/wf/api"

	"wf_api/server/internal/svc"
	"wf_api/server/internal/types"

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
	return internal.Success("", append(api.EventShops(), api.BossShops()...))
}
