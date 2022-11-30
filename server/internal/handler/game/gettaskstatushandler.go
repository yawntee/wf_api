package game

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wf_api/server/internal/logic/game"
	"wf_api/server/internal/svc"
)

func GetTaskStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := game.NewGetTaskStatusLogic(r.Context(), svcCtx)
		resp, err := l.GetTaskStatus()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
