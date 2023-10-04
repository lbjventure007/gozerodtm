package handler

import (
	"net/http"

	"awesomeProject7/users/userapi/internal/logic"
	"awesomeProject7/users/userapi/internal/svc"
	"awesomeProject7/users/userapi/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TranInOutBalanceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TranInOutRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewTranInOutBalanceLogic(r.Context(), svcCtx)
		resp, err := l.TranInOutBalance(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
