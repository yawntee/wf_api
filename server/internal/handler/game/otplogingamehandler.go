package game

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wf_api/server/internal/logic/game"
	"wf_api/server/internal/svc"
	"wf_api/server/internal/types"
)

func OtpLoginGameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OtpLoginGameReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := game.NewOtpLoginGameLogic(r.Context(), svcCtx)
		resp, err := l.OtpLoginGame(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
