package action

import (
	"emq_plugin_server/emqWeb/Infra/enum"
	"emq_plugin_server/emqWeb/domain/model"
	"emq_plugin_server/emqWeb/domain/service"
	"fmt"
)

// 登录操作
func LoginH(req model.RequestJsonModel) model.ParamModel {
	// do something(不关乎领域逻辑和业务的)
	fmt.Println("-------")
	// 调用domain 内的EMQ领域服务
	service.UserLogin(req)
	return model.ParamModel{ErrorCode: enum.OPERATE_SUCCESS}
}
