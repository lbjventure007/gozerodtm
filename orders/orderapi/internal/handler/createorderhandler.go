package handler

import (
	"awesomeProject7/orders/orderapi/internal/logic"
	"awesomeProject7/orders/orderapi/internal/svc"
	"awesomeProject7/orders/orderapi/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func CreateOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrderRequest

		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateOrderLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
