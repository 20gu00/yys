package homestayOrder

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-travel/service/order/cmd/api/internal/logic/homestayOrder"
	"go-travel/service/order/cmd/api/internal/svc"
	"go-travel/service/order/cmd/api/internal/types"
)

func UserHomestayOrderDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserHomestayOrderDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := homestayOrder.NewUserHomestayOrderDetailLogic(r.Context(), svcCtx)
		resp, err := l.UserHomestayOrderDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
