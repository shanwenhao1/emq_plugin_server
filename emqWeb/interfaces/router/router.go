package router

import (
	"emq_plugin_server/emqWeb/app/action"
	"github.com/gin-gonic/gin"
)

// 路由映射, 为用户请求服务入口, 映射至application层中服务
func Router(handleMap map[string]gin.HandlerFunc) {
	userAction := new(action.UserJsonModel)
	emqAcl := new(action.EmqAclModel)
	emqWebHook := new(action.EmqWebHookModel)

	// EMQ测试鉴权所用http_auth_acl
	handleMap["mqtt/auth"] = emqAcl.AuthAcl
	handleMap["mqtt/superuser"] = emqAcl.SuperAcl
	handleMap["mqtt/acl"] = emqAcl.AclReq

	// EMQ web hook 测试所用
	handleMap["mqtt/webhook"] = emqWebHook.WebHook

	// 添加示例login路由
	handleMap["login"] = userAction.Login
}
