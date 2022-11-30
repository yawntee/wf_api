package game

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wf_api/server/internal/logic/game"
	"wf_api/server/internal/svc"
	"wf_api/server/internal/types"
)

func SendLoginGameOtpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendOtpReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := game.NewSendLoginGameOtpLogic(r.Context(), svcCtx)
		resp, err := l.SendLoginGameOtp(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
