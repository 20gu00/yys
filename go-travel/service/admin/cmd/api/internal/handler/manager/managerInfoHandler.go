package manager

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-travel/service/admin/cmd/api/internal/logic/manager"
	"go-travel/service/admin/cmd/api/internal/svc"
)

func ManagerInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := manager.NewManagerInfoLogic(r.Context(), svcCtx)
		resp, err := l.ManagerInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
