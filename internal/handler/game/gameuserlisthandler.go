package game

import (
	"net/http"
	"wf_api/internal/logic/game"
	"wf_api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GameUserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := game.NewGameUserListLogic(r.Context(), svcCtx)
		resp, err := l.GameUserList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
