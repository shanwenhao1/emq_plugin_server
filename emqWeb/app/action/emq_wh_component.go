package action

import (
	"emq_plugin_server/emqWeb/Infra/enum"
	"emq_plugin_server/emqWeb/domain/model"
	"fmt"
)

// Web Hook钩子操作
func WebHookFunc(req interface{}) model.ParamModel {
	fmt.Println("===========---sss", req)
	var result model.ParamModel
	switch test := req.(type) {
	case model.WHClientConn:
		fmt.Println("--------------------------1")
		act := req.(model.WHClientConn)
		// 调用领域内逻辑处理函数
		result = act.ActionMethod()
	case model.WHClientDisConn:
		fmt.Println("--------------------------2")
		act := req.(model.WHClientDisConn)
		result = act.ActionMethod()
	case model.WHClientSub:
		fmt.Println("--------------------------3")
		act := req.(model.WHClientSub)
		result = act.ActionMethod()
	case model.WHClientUnSub:
		fmt.Println("--------------------------4")
		act := req.(model.WHClientUnSub)
		result = act.ActionMethod()
	case model.WHSessionCreate:
		fmt.Println("--------------------------5")
		act := req.(model.WHSessionCreate)
		result = act.ActionMethod()
	case model.WHSessionSub:
		fmt.Println("--------------------------6")
		act := req.(model.WHSessionSub)
		result = act.ActionMethod()
	case model.WHSessionUnSub:
		fmt.Println("--------------------------7")
		act := req.(model.WHSessionUnSub)
		result = act.ActionMethod()
	case model.WHSessionT:
		fmt.Println("--------------------------8")
		act := req.(model.WHSessionT)
		result = act.ActionMethod()
	case model.WHMsgPub:
		fmt.Println("--------------------------9")
		act := req.(model.WHMsgPub)
		result = act.ActionMethod()
	case model.WHMsgDeliver:
		fmt.Println("--------------------------10")
		act := req.(model.WHMsgDeliver)
		result = act.ActionMethod()
	case model.WHMsgAck:
		fmt.Println("--------------------------11")
		act := req.(model.WHMsgAck)
		result = act.ActionMethod()
	default:
		fmt.Println("--------------------------12", test)
		result = model.ParamModel{ErrorCode: enum.ACTION_FAILED}
	}
	return result
}
