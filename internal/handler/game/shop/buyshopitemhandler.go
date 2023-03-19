package shop

import (
	"net/http"
	"wf_api/internal/logic/game/shop"
	"wf_api/internal/svc"
	"wf_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func BuyShopItemHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BuyShopItemReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := shop.NewBuyShopItemLogic(r.Context(), svcCtx)
		resp, err := l.BuyShopItem(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
