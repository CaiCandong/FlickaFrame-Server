// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/oss/endpoint",
				Handler: EndpointHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/oss/uptoken",
				Handler: CreateUpTokenHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)
}