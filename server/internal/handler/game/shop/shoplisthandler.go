package shop

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wf_api/server/internal/logic/game/shop"
	"wf_api/server/internal/svc"
)

func ShopListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := shop.NewShopListLogic(r.Context(), svcCtx)
		resp, err := l.ShopList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
