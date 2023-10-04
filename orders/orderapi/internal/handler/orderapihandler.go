package handler

import (
	"net/http"

	"awesomeProject7/orders/orderapi/internal/logic"
	"awesomeProject7/orders/orderapi/internal/svc"
	"awesomeProject7/orders/orderapi/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderapiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOrderapiLogic(r.Context(), svcCtx)
		resp, err := l.Orderapi(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
