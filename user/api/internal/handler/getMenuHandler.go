package handler

import (
	"devops/common/response"
	"devops/user/api/internal/logic"
	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMenuRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJson(w, 200, response.HandlerError(err))
			return
		}

		l := logic.NewGetMenuLogic(r.Context(), svcCtx)
		resp, err := l.GetMenu(&req)
		if err != nil {
			httpx.WriteJson(w, 200, response.HandlerError(err))
		} else {
			httpx.OkJson(w, response.HandlerResp(resp))
		}
	}
}