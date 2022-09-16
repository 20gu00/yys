package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-travel/service/usercenter/cmd/api/internal/logic/user"
	"go-travel/service/usercenter/cmd/api/internal/svc"
	"go-travel/service/usercenter/cmd/api/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
