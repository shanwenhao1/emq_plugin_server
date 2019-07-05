package action

import (
	"emq_plugin_server/emqWeb/domain/model"
	"emq_plugin_server/emqWeb/domain/service"
	"fmt"
)

// Auth认证操作
func AuthAcl(req model.EmqAcl) model.ParamModel {
	// do something(不关乎领域逻辑和业务的)
	fmt.Println("-------", req)
	// 调用domain 内的User领域服务
	result := service.AuthAcl(req)
	return result
}

// Super user 认证操作
func SuperAcl(req model.EmqAcl) model.ParamModel {
	result := service.SuperAcl(req)
	return result
}

// Super user 认证操作
func AclReq(req model.EmqAcl) model.ParamModel {
	result := service.AclReq(req)
	return result
}
