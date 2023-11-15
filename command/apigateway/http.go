package apigateway

import (
	"gtstart/config"
	"gtstart/design/gen/types"
	"gtstart/internal/handler"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandlers(server *rest.Server) {
	// 需要jwt认证的
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: TestHandler(),
			},
		}, rest.WithJwt(config.C.Jwt.Secret))
	// 不需要jwt的
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from_no/:name",
				Handler: TestHandler(),
			},
		})
}

func TestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := handler.NewGreetLogic(logx.WithContext(r.Context()))
		resp, err := l.Greet(r.Context(), &req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
		return
	}
}
