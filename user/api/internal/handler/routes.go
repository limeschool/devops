// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"devops/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/check_health",
				Handler: CheckHealthHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/changepwd",
				Handler: ChangePasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/auth/validate",
				Handler: AuthHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/info",
				Handler: GetUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/page",
				Handler: GetUserPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/info",
				Handler: AddUserHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/user/info",
				Handler: DeleteUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/user/info",
				Handler: UpdateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/role/list",
				Handler: GetRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/role/info",
				Handler: AddRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/role/info",
				Handler: DeleteRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/role/info",
				Handler: UpdateRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/team/list",
				Handler: GetTeamHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/team/info",
				Handler: AddTeamHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/team/info",
				Handler: DeleteTeamHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/team/info",
				Handler: UpdateTeamHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/menu/list",
				Handler: GetMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/menu/info",
				Handler: AddMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/menu/info",
				Handler: DeleteMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/menu/info",
				Handler: UpdateMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/role/menu",
				Handler: CreateRoleMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/role/menu_ids",
				Handler: GetRoleMenuIdsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/role/menus",
				Handler: GetRoleMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/login/log",
				Handler: GetLogHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/ums"),
	)
}
