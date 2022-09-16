package homestay

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"go-travel/service/admin/cmd/api/internal/logic/homestay"
	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"
)

func HomestayActivityAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddHomestayActivityReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := homestay.NewHomestayActivityAddLogic(r.Context(), svcCtx)
		resp, err := l.HomestayActivityAdd(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
