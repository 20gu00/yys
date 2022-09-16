package homestayBussiness

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-travel/service/travel/cmd/api/internal/logic/homestayBussiness"
	"go-travel/service/travel/cmd/api/internal/svc"
	"go-travel/service/travel/cmd/api/internal/types"
)

func HomestayBussinessListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayBussinessListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := homestayBussiness.NewHomestayBussinessListLogic(r.Context(), svcCtx)
		resp, err := l.HomestayBussinessList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
