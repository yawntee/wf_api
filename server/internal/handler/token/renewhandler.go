package token

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wf_api/server/internal/logic/token"
	"wf_api/server/internal/svc"
)

func RenewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := token.NewRenewLogic(r.Context(), svcCtx)
		resp, err := l.Renew()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
