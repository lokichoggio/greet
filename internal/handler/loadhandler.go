package handler

import (
	"net/http"

	"greet/internal/logic"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoadHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoadLogic(r.Context(), ctx)
		resp, err := l.Load(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
