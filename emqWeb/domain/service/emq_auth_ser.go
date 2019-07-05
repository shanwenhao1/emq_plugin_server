package service

import (
	"emq_plugin_server/emqWeb/Infra/enum"
	"emq_plugin_server/emqWeb/domain/model"
)

// emq auth acl 逻辑处理
func AuthAcl(req model.EmqAcl) model.ParamModel {
	// do something
	return model.ParamModel{ErrorCode: enum.ACTION_SUCCESS}
}

// emq super user acl 逻辑处理
func SuperAcl(req model.EmqAcl) model.ParamModel {
	// do something
	return model.ParamModel{ErrorCode: enum.ACTION_SUCCESS}
}

// emq action acl 逻辑处理
func AclReq(req model.EmqAcl) model.ParamModel {
	// do something
	return model.ParamModel{ErrorCode: enum.ACTION_SUCCESS}
}
